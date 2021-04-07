Workspace is a directory containing a hierarchy of directories in which files related to the golang project are stored. <br>
The advantage of this is a common organization of files => Ease of sharing. <br>
The workspace directory is defined in the `GOPATH` environment variable which is usually set during installation. <br>
The Go tools assume that all the code to be executed is in the location specified in `GOPATH`. <br>
Generally a single workspace is used for multiple projects. <br>
  
## Default workspace directory layout
The default workspace layout consists of the following:

### src
Directory containing source code files.
  
### pkg
Directory containing packages/libraries.
  
### bin
Directory containing executables.
  
<br>

**NOTE**: Directory hierarchies are recommended NOT enforced, so we can also have an executable file in `src` if we want. <br>
