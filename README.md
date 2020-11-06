# awstool

Inteview simple tool

# Build

`make build`

# Usage

To use this simple tool first we need to change some variables in cmd/main.go file:

```
const (
	arnTargetgroup = "arn:aws:elasticloadbalancing:eu-central-1:XXXXXXXXXXXX:targetgroup/test/c7a44c37a781fe94"
	region         = "eu-central-1"
)
```

We need to setup arnTargetgroup and region for our AWS. arnTargetgroup we can check by executing command:

`aws elbv2 --region=<region> describe-target-groups`

Then we can use our application this way:
```
./deploy --help
Usage of ./deploy:
  -firstImage string
        image which we want to change (default "ami-0c960b947cbb2dd16")
  -secondImage string
        image to which we want to change (default "ami-0ed0be684d3f014bf")
```