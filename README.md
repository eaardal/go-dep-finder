# Go dep finder

Scan Go projects to list (direct) usages of a specified dependency.

All code written by ChatGPT 4.

### Example

Assuming the Go projects are next to each other under a root folder:

```
my_work_projects
    /go_project_a
        go.mod
    /go_project_b
        go.mod
    /go_project_c
        go.mod
```

Run go-dep-finder against the root folder and a dependency you want to find usages of:

```shell
go-dep-finder /path/to/my_work_projects some_lib_name

Project: go_project_a, Depedency Version: v0.4.11
Project: go_project_b, Depedency Version: v0.12.3
Project: go_project_c, Depedency Version: v0.7.8
```