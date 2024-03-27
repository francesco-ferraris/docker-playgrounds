## Usage

```makefile
make
```
Same as `make run` but it enables the watching configuration defined in the `compose.yaml` file.
<br/>
<br/>

```makefile
make run
```
Run the containers using docker compose. It automatically build and run all the containers defined in the `compose.yaml` file.
<br/>
<br/>

```makefile
make clean
```
Stop and delete the containers created with the `make run` directive.
