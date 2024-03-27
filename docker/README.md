## Usage

```makefile
make
```
Default action that build and run the application in a container.
<br/>
<br/>

```makefile
make build
```
Build the docker image using the Dockerfile in the folder
<br/>
<br/>
```makefile
make run
```
Run the container. Before executing this directive, 
the container image must be built with the `make build` directive at least one time.
Running without building the image after modifying the source code results in
running a container without the latest updates.
<br/>
<br/>

```makefile
make clean
```
Delete the image from the local system.
