package ansible

var (
	ansibleTemplate = `{{ .Shebang }}
# -*- coding: utf-8 -*-

_ANSIBALLZ_WRAPPER = True # For test-module.py script to tell this is a ANSIBALLZ_WRAPPER
# This code is part of Ansible, but is an independent component.
# The code in this particular templatable string, and this templatable string
# only, is BSD licensed.  Modules which end up using this snippet, which is
# dynamically combined together by Ansible still belong to the author of the
# module, and they may assign their own license to the complete work.
#
# Copyright (c), James Cammarata, 2016
# Copyright (c), Toshio Kuratomi, 2016
#
# Redistribution and use in source and binary forms, with or without modification,
# are permitted provided that the following conditions are met:
#
#    * Redistributions of source code must retain the above copyright
#      notice, this list of conditions and the following disclaimer.
#    * Redistributions in binary form must reproduce the above copyright notice,
#      this list of conditions and the following disclaimer in the documentation
#      and/or other materials provided with the distribution.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
# ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
# WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
# IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
# INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
# PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
# INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
# LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE
# USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
def _ansiballz_main():
    import os
    import os.path
    # Access to the working directory is required by Python when using pipelining, as well as for the coverage module.
    # Some platforms, such as macOS, may not allow querying the working directory when using become to drop privileges.
    try:
        os.getcwd()
    except OSError:
        try:
            os.chdir(os.path.expanduser('~'))
        except OSError:
            os.chdir('/')

    import sys
    import __main__
    # For some distros and python versions we pick up this script in the temporary
    # directory.  This leads to problems when the ansible module masks a python
    # library that another import needs.  We have not figured out what about the
    # specific distros and python versions causes this to behave differently.
    #
    # Tested distros:
    # Fedora23 with python3.4  Works
    # Ubuntu15.10 with python2.7  Works
    # Ubuntu15.10 with python3.4  Fails without this
    # Ubuntu16.04.1 with python3.5  Fails without this
    # To test on another platform:
    # * use the copy module (since this shadows the stdlib copy module)
    # * Turn off pipelining
    # * Make sure that the destination file does not exist
    # * ansible ubuntu16-test -m copy -a 'src=/etc/motd dest=/var/tmp/m'
    # This will traceback in shutil.  Looking at the complete traceback will show
    # that shutil is importing copy which finds the ansible module instead of the
    # stdlib module
    scriptdir = None
    try:
        scriptdir = os.path.dirname(os.path.realpath(__main__.__file__))
    except (AttributeError, OSError):
        # Some platforms don't set __file__ when reading from stdin
        # OSX raises OSError if using abspath() in a directory we don't have
        # permission to read (realpath calls abspath)
        pass
    # Strip cwd from sys.path to avoid potential permissions issues
    excludes = set(('', '.', scriptdir))
    sys.path = [p for p in sys.path if p not in excludes]
    import base64
    import runpy
    import shutil
    import tempfile
    import zipfile
    if sys.version_info < (3,):
        PY3 = False
    else:
        PY3 = True
    
    ZIPDATA = """{{ .ZipData }}"""

    # Note: temp_path isn't needed once we switch to zipimport
    def invoke_module(modlib_path, temp_path, json_params):
        # When installed via setuptools (including python setup.py install),
        # ansible may be installed with an easy-install.pth file.  That file
        # may load the system-wide install of ansible rather than the one in
        # the module.  sitecustomize is the only way to override that setting.
        z = zipfile.ZipFile(modlib_path, mode='a')
        # py3: modlib_path will be text, py2: it's bytes.  Need bytes at the end
        sitecustomize = u'import sys\\nsys.path.insert(0,"%s")\\n' %  modlib_path
        sitecustomize = sitecustomize.encode('utf-8')
        # Use a ZipInfo to work around zipfile limitation on hosts with
        # clocks set to a pre-1980 year (for instance, Raspberry Pi)
        zinfo = zipfile.ZipInfo()
        zinfo.filename = 'sitecustomize.py'
        zinfo.date_time = ( {{ .Year }}, {{ .Month }}, {{ .Day }}, {{ .Hour }}, {{ .Minute }}, {{ .Second }} )
        z.writestr(zinfo, sitecustomize)
        z.close()
        # Put the zipped up module_utils we got from the controller first in the python path so that we
        # can monkeypatch the right basic
        sys.path.insert(0, modlib_path)
        # Monkeypatch the parameters into basic
        from ansible.module_utils import basic
        basic._ANSIBLE_ARGS = json_params

        # Run the module!  By importing it as '__main__', it thinks it is executing as a script
        runpy.run_module(mod_name='{{ .ModuleFqn }}', init_globals=dict(_module_fqn='{{ .ModuleFqn }}', _modlib_path=modlib_path),
                         run_name='__main__', alter_sys=True)
        # Ansible modules must exit themselves
        print('{"msg": "New-style module did not handle its own exit", "failed": true}')
        sys.exit(1)
    def debug(command, zipped_mod, json_params):
        # The code here normally doesn't run.  It's only used for debugging on the
        # remote machine.
        #
        # The subcommands in this function make it easier to debug ansiballz
        # modules.  Here's the basic steps:
        #
        # Run ansible with the environment variable: ANSIBLE_KEEP_REMOTE_FILES=1 and -vvv
        # to save the module file remotely::
        #   $ ANSIBLE_KEEP_REMOTE_FILES=1 ansible host1 -m ping -a 'data=october' -vvv
        #
        # Part of the verbose output will tell you where on the remote machine the
        # module was written to::
        #   [...]
        #   <host1> SSH: EXEC ssh -C -q -o ControlMaster=auto -o ControlPersist=60s -o KbdInteractiveAuthentication=no -o
        #   PreferredAuthentications=gssapi-with-mic,gssapi-keyex,hostbased,publickey -o PasswordAuthentication=no -o ConnectTimeout=10 -o
        #   ControlPath=/home/badger/.ansible/cp/ansible-ssh-%h-%p-%r -tt rhel7 '/bin/sh -c '"'"'LANG=en_US.UTF-8 LC_ALL=en_US.UTF-8
        #   LC_MESSAGES=en_US.UTF-8 /usr/bin/python /home/badger/.ansible/tmp/ansible-tmp-1461173013.93-9076457629738/ping'"'"''
        #   [...]
        #
        # Login to the remote machine and run the module file via from the previous
        # step with the explode subcommand to extract the module payload into
        # source files::
        #   $ ssh host1
        #   $ /usr/bin/python /home/badger/.ansible/tmp/ansible-tmp-1461173013.93-9076457629738/ping explode
        #   Module expanded into:
        #   /home/badger/.ansible/tmp/ansible-tmp-1461173408.08-279692652635227/ansible
        #
        # You can now edit the source files to instrument the code or experiment with
        # different parameter values.  When you're ready to run the code you've modified
        # (instead of the code from the actual zipped module), use the execute subcommand like this::
        #   $ /usr/bin/python /home/badger/.ansible/tmp/ansible-tmp-1461173013.93-9076457629738/ping execute
        # Okay to use __file__ here because we're running from a kept file
        basedir = os.path.join(os.path.abspath(os.path.dirname(__file__)), 'debug_dir')
        args_path = os.path.join(basedir, 'args')
        if command == 'explode':
            # transform the ZIPDATA into an exploded directory of code and then
            # print the path to the code.  This is an easy way for people to look
            # at the code on the remote machine for debugging it in that
            # environment
            z = zipfile.ZipFile(zipped_mod)
            for filename in z.namelist():
                if filename.startswith('/'):
                    raise Exception('Something wrong with this module zip file: should not contain absolute paths')
                dest_filename = os.path.join(basedir, filename)
                if dest_filename.endswith(os.path.sep) and not os.path.exists(dest_filename):
                    os.makedirs(dest_filename)
                else:
                    directory = os.path.dirname(dest_filename)
                    if not os.path.exists(directory):
                        os.makedirs(directory)
                    f = open(dest_filename, 'wb')
                    f.write(z.read(filename))
                    f.close()
            # write the args file
            f = open(args_path, 'wb')
            f.write(json_params)
            f.close()
            print('Module expanded into:')
            print('%s' % basedir)
            exitcode = 0
        elif command == 'execute':
            # Execute the exploded code instead of executing the module from the
            # embedded ZIPDATA.  This allows people to easily run their modified
            # code on the remote machine to see how changes will affect it.
            # Set pythonpath to the debug dir
            sys.path.insert(0, basedir)
            # read in the args file which the user may have modified
            with open(args_path, 'rb') as f:
                json_params = f.read()
            # Monkeypatch the parameters into basic
            from ansible.module_utils import basic
            basic._ANSIBLE_ARGS = json_params
            # Run the module!  By importing it as '__main__', it thinks it is executing as a script
            runpy.run_module(mod_name='{{ .ModuleFqn }}', init_globals=None, run_name='__main__', alter_sys=True)
            # Ansible modules must exit themselves
            print('{"msg": "New-style module did not handle its own exit", "failed": true}')
            sys.exit(1)
        else:
            print('WARNING: Unknown debug command.  Doing nothing.')
            exitcode = 0
        return exitcode
    #
    # See comments in the debug() method for information on debugging
    #
    ANSIBALLZ_PARAMS = '{{ .Params }}'

    if PY3:
        ANSIBALLZ_PARAMS = ANSIBALLZ_PARAMS.encode('utf-8')
    try:
        # There's a race condition with the controller removing the
        # remote_tmpdir and this module executing under async.  So we cannot
        # store this in remote_tmpdir (use system tempdir instead)
        # Only need to use [ansible_module]_payload_ in the temp_path until we move to zipimport
        # (this helps ansible-test produce coverage stats)
        temp_path = tempfile.mkdtemp(prefix='ansible_{{ .AnsibleModule }}_payload_')
        zipped_mod = os.path.join(temp_path, 'ansible_{{ .AnsibleModule }}_payload.zip')
        with open(zipped_mod, 'wb') as modlib:
            modlib.write(base64.b64decode(ZIPDATA))
        if len(sys.argv) == 2:
            exitcode = debug(sys.argv[1], zipped_mod, ANSIBALLZ_PARAMS)
        else:
            # Note: temp_path isn't needed once we switch to zipimport
            invoke_module(zipped_mod, temp_path, ANSIBALLZ_PARAMS)
    finally:
        try:
            shutil.rmtree(temp_path)
        except (NameError, OSError):
            # tempdir creation probably failed
            pass
    sys.exit(exitcode)
if __name__ == '__main__':
    _ansiballz_main()
	`
)
