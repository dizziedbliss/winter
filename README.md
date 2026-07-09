### For v0.1.0

July 12, 2026 

In cmd/cli

- I've to do thing that reads command from cli, like `winter deploy`
- Which also parse the cmd after recognising it , like the function `deploy`

In deployment module

- Deploy.go will have pipeline for deployment fns from other module such as docker if the deploy has a dockerfile so maybe for discovery & validation use another file from the same module. ‘cause deploy might have more than docker later on .
- Acc what file the discovery finds n validates for now dockerfile it'll run that module later like docker.run

Docker module

- It'll have run build stop logs n others files which does their work upon calling.

We're basically doing

`Docker build abc`

`Docker run abc`

`Docker logs abc`
 But with our command winter 😭😭🙏🏻

July 9, 2026 

### Done with v0.2.0

done with  CI but with alot of errors tho hehe…

```jsx
|   |-- deployment
|   |   |-- build.go
|   |   |-- deploy.go
|   |   |-- scan.go
|   |   `-- validate.go
```

#### TODOs:

- [ ]  super clean logs with error reasoning
- [ ]  more flags for each edge cases
- EdgeCases
    - Scan
        - multiple dockerfiles
        - permissions
        - hidden dirs
        - .git n .gitignores
        - symlinks
    - Validate
        - ofc each of them should be validated
    - build
        - no docker installed lol
        - no daemon running
        - offline network
        - missing files n etc
        - timeout

### For v0.3.0

make CD, for local deployment, and add more config options, ie., winter should understand the project structure and maybe use its own config instead of commands and alot of  flags hehe just like how docker does and gorelease works