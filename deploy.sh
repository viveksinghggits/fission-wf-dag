set +x
fission environment create  --name go --image fission/go-env --builder fission/go-builder
fission fn create  --name first --env go --src first.go --entrypoint Handler
fission fn create  --name second --env go --src second.go --entrypoint Handler
fission fn create  --name third --env go --src third.go --entrypoint Handler
fission fn create  --name fourth --env go --src fourth.go --entrypoint Handler
fission fn create  --name fifth --env go --src fifth.go --entrypoint Handler
fission fn create  --name sixth --env go --src sixth.go --entrypoint Handler
fission function create  --name complexwf --env workflow --src complex-wf.yaml
#fission route create --method GET --url /complexwf --function complexwf


#fission fn update  --name first --env go --src first.go --entrypoint Handler
#fission fn update  --name second --env go --src second.go --entrypoint Handler
#fission fn update  --name third --env go --src third.go --entrypoint Handler
#fission fn update  --name fourth --env go --src fourth.go --entrypoint Handler
#fission fn update  --name fifth --env go --src fifth.go --entrypoint Handler
#fission fn update  --name sixth --env go --src sixth.go --entrypoint Handler

#fission function create --name complexwf --env workflow --src complex-wf.yaml --fnNamespace fission --envNamespace fission
