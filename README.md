# team-env

Store environment variables remotely and share them with your team.

Currently only the `ssm` remote store is supported, but the package can be easily extended for adding more stores in the future.

- `ssm`: Store variables in your AWS account, using the [SSM Parameters Store](https://docs.aws.amazon.com/es_es/systems-manager/latest/userguide/systems-manager-paramstore.html).

## Installation:

```bash
go get -u github.com/andybar2/team-env
```

## Configuration:

Create the `team-env.json` file in your project folder, and set your configuration parameters there. Here is an example using the `ssm` store on AWS.

```json
{
  "store": "ssm",
  "project": "example",
  "aws_profile": "myaccount",
  "aws_region": "us-east-2"
}
```

The value of `aws_profile` is a reference to your `~/.aws/credentials` file. The keys for the given profile will be used by `team-env` to connect to your AWS account:

```
[myaccount]
aws_access_key_id = xxxxxxxx
aws_secret_access_key = xxxxxxxxxxxxxxxxxxxxxxxx
```

## Usage:

### Set variable:

```bash
team-env set --env=development --var=DEBUG --val=1
```

### Get variable:

```bash
team-env get --env=development --var=DEBUG
```

### Delete variable:

```bash
team-env del --env=development --var=DEBUG
```

### Print all variables for an environment:

```bash
team-env print --env=development
```

### Configure your local environment with the values of the variables in the remote store:

```bash
export $(team-env print --env=development)
```
