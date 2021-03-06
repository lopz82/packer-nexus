# Packer template for Nexus Repository OSS

## About The Project

This is an unofficial Packer template to create a Nexus Repository OSS AMI.


## Prerequisites

You will need the following tools:

* An [AWS](https://aws.amazon.com/?nc2=h_lg) account
* [Packer](https://learn.hashicorp.com/tutorials/packer/getting-started-install)

* An AWS Security Group with access to port 22
* An AWS SSH key pair

Optional, just for testing:

* [Terraform](https://www.terraform.io/downloads.html)
* [Go](https://golang.org/doc/install)

## Usage

### Configuration

You can set up some configurations in variables.json

* `version`: Nexus Repository Manager OSS version to be installed. You can check older versions [here](https://help.sonatype.com/repomanager3/download/download-archives---repository-manager-3)
* `data-path`: path where Nexus Repository will store all its data. You can configure it to point to a dir where you want to mount an EBS, for example.
* `ubuntu_version`: Ubuntu version used as base image. (More images)[https://cloud-images.ubuntu.com/locator/ec2/]

```json
{
  "version": "3.29.0-02",
  "data-path": "/datadir/nexus",
  "ubuntu_version": "ubuntu-bionic-18.04"
}
```

### Building the image

We will build the AMI using packer.

```shell
$ git clone https://github.com/lopz82/packer-nexus.git
$ cd packer-nexus
$ packer build -var-file=nexus/variables.json nexus/packer.json
```

Once Packer is done, you should see a new AMI in the EC2 section of your AWS console (Images/AMIs).

> Note: for convenience, a manifests.json file is created. This file can be used to test the image.

## Testing
The tests are written in go and use [Terratest](https://terratest.gruntwork.io/).
### Terraform configuration

In order to test the image, you will need to configure the following:

* AWS Key Pair
* Security Group granting access to port 22

> Make sure you add the public key to your ssh agent `ssh-add ~/.ssh/YourKey.pem`


### Testing the image

Just run the following commands:

```shell
$ cd test
$ go test -v -timeout 30m -ami $(cat ../nexus/manifest.json | jq -r '.builds[0].artifact_id' | grep -o 'ami.*$') ./...
```

Double check that the test instance is terminated. That should happen even if the tests fail, 
but it is recommended not to incur unexpected expenses.

> You can also provide manually the ami instead of extracting it from `manifests.json`.

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
