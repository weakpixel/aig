package ansible

var (
	ansibleTemplateRemote = `{{ .Shebang }}
# -*- coding: utf-8 -*-

_ANSIBALLZ_WRAPPER = True # For test-module.py script to tell this is a ANSIBALLZ_WRAPPER
def _ansiballz_main():
    import os
    import os.path
    try:
        os.getcwd()
    except OSError:
        try:
            os.chdir(os.path.expanduser('~'))
        except OSError:
            os.chdir('/')

    import sys
    import __main__
    scriptdir = None
    try:
        scriptdir = os.path.dirname(os.path.realpath(__main__.__file__))
    except (AttributeError, OSError):
        pass
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
    
    def invoke_module(modlib_path, temp_path, json_params):
        z = zipfile.ZipFile(modlib_path, mode='a')
        sitecustomize = u'import sys\\nsys.path.insert(0,"%s")\\n' %  modlib_path
        sitecustomize = sitecustomize.encode('utf-8')
        zinfo = zipfile.ZipInfo()
        zinfo.filename = 'sitecustomize.py'
        zinfo.date_time = ( {{ .Year }}, {{ .Month }}, {{ .Day }}, {{ .Hour }}, {{ .Minute }}, {{ .Second }} )
        z.writestr(zinfo, sitecustomize)
        z.close()
        sys.path.insert(0, modlib_path)
        from ansible.module_utils import basic
        basic._ANSIBLE_ARGS = json_params

        runpy.run_module(mod_name='{{ .ModuleFqn }}', init_globals=dict(_module_fqn='{{ .ModuleFqn }}', _modlib_path=modlib_path),
                         run_name='__main__', alter_sys=True)
        print('{"msg": "New-style module did not handle its own exit", "failed": true}')
        sys.exit(1)
    
    
    ANSIBALLZ_PARAMS = '{{ .Params }}'

    temp_path = tempfile.mkdtemp(prefix='ansible_{{ .AnsibleModule }}_payload_')
    
    if PY3:
        ANSIBALLZ_PARAMS = ANSIBALLZ_PARAMS.encode('utf-8')
    try:
        invoke_module("/tmp/ansible_modules.zip", temp_path, ANSIBALLZ_PARAMS)
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
