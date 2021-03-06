# team

Save your project configuration (environment variables, configuration files, etc) in a remote store and easily share it with your team instead of having to store it inside your project repository.

This allows you to securely share the source code of your project without including secrets, API keys, service account files, etc.

Currently only the `aws` remote store is supported, but the package can be extended for adding more stores in the future.

- `aws`: Stores project configuration in your AWS account. Environment variables are stored in the [SSM Parameters Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-paramstore.html), and files are stored in [S3](https://docs.aws.amazon.com/s3/index.html).

## Installation:

```bash
go get -u github.com/andybar2/team
```

## Configuration:

Create the `team.json` file in your project folder, and set your configuration parameters there.

```json
{
  "project": "example",
  "store": "aws",
  "aws_profile": "myaccount",
  "aws_region": "us-east-2"
}
```

The value of `aws_profile` is a reference to your `~/.aws/credentials` file. The keys for the given profile will be used by `team` to connect to your AWS account:

```
[myaccount]
aws_access_key_id = xxxxxxxx
aws_secret_access_key = xxxxxxxxxxxxxxxxxxxxxxxx
```

## Manage environment variables:

### Set variable:

```bash
team env set -s "development" -n "STRIPE_API_KEY" -v "1234567890"
```

### Get variable:

```bash
team env get -s "development" -n "STRIPE_API_KEY"
```

### Delete variable:

```bash
team env delete -s "development" -n "STRIPE_API_KEY"
```

### Print all variables for a stage:

```bash
team env print -s "development"
```

### Configure your local environment with all the variables in a stage:

```bash
export $(team env print -s "development")
```

### Import all the environment variables on a stage to a local file:

```bash
mkdir -p .team/development
team env print -s "development" > .team/development/env
```

**Tip**: Add `.team` to your `.gitignore` file :)

## Manage configuration files:

### Upload file:

```bash
team files upload -s "development" -p ".team/development/google-service-account.json"
```

### Download file:

```bash
team files download -s "development" -p ".team/development/google-service-account.json"
```

### Delete file:

```bash
team files delete -s "development" -p ".team/development/google-service-account.json"
```

### List all files in a stage:

```bash
team files list -s "development"
```
