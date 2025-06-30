```
NAME:
   justmcp - An MCP Server for Just

USAGE:
   justmcp [global options] [command [command options]]

VERSION:
   0.0.2

COMMANDS:
   dump     dump justfile for debugging
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --[no-]chdir                                                             cd to the same file directory as justfile when running (default: true) [$JUSTMCP_CHDIR]
   --[no-]mise                                                              use 'mise x' when running just recipes (default: false) [$JUSTMCP_MISE]
   --justfile justfile, -f justfile                                         path to justfile [$JUSTMCP_JUSTFILE, $JUST_JUSTFILE]
   --[no-]minimal                                                           only register minimal tools (default: false) [$JUSTMCP_MINIMAL]
   --recipes recipe(s), -r recipe(s) [ --recipes recipe(s), -r recipe(s) ]  add the given recipe(s) as tool(s) [$JUSTMCP_RECIPES]
   --help, -h                                                               show help
   --version, -v                                                            print the version
```
