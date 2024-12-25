# Amazon Elastic Kubernetes Service (EKS)

Amazon Elastic Kubernetes Service (EKS) is a managed Kubernetes service provided by AWS that simplifies the deployment, management, and scaling of containerized applications. With EKS, you don't need to install or operate your own Kubernetes control plane, allowing you to focus on application development.

## About Amazon EKS
Amazon EKS integrates seamlessly with AWS services such as IAM, VPC, and CloudWatch, ensuring a secure and scalable environment for your Kubernetes workloads. It supports hybrid deployments, enabling applications to run on both the cloud and on-premises infrastructure.

## Key Features
- Highly available and secure Kubernetes control plane.
- Integration with AWS services like IAM, VPC, and CloudWatch.
- Support for hybrid workloads with EKS Anywhere.
- Seamless Kubernetes version upgrades and updates.

## Benefits
- Reduces operational overhead with managed Kubernetes control plane.
- Ensures reliability, scalability, and security for your applications.
- Provides flexibility to run workloads on the cloud or on-premises.

## How to Use
1. **Set Up Your EKS Cluster:**
   - Use the AWS Management Console, AWS CLI, or eksctl to create and configure your cluster.
2. **Deploy Your Applications:**
   - Use Kubernetes manifests or Helm charts to deploy applications to your EKS cluster.
3. **Monitor and Scale:**
   - Integrate with Amazon CloudWatch for monitoring and leverage Kubernetes autoscaling capabilities.

## Learn More
For detailed information and step-by-step guides, visit the [Amazon EKS Documentation](https://aws.amazon.com/eks/).

---

# Docker Compose

Docker Compose is a tool for defining and managing multi-container Docker applications. With Compose, you can use a YAML file to configure your application's services, making it easy to start, stop, and scale your application using a single command.

## About Docker Compose
Docker Compose simplifies the development and deployment of applications that require multiple containers, such as a web application with a frontend, backend, and database. By defining all the services in a single file, Docker Compose enables developers to manage the entire application stack efficiently.

## Key Features
- Define multi-container applications in a single `docker-compose.yml` file.
- Start, stop, and scale services with a single command.
- Support for variable substitution and configuration templates.
- Networking between containers is automatically configured.

## Benefits
- Simplifies the management of complex applications with multiple services.
- Provides consistency across development, testing, and production environments.
- Saves time by automating service orchestration.

## How to Use
1. **Install Docker Compose:**
   - Ensure Docker is installed, then install Docker Compose following the [official documentation](https://docs.docker.com/compose/install/).
2. **Create a `docker-compose.yml` File:**
   - Define your services, networks, and volumes in the YAML file.
3. **Run Your Application:**
   - Use the command `docker-compose up` to start all services defined in the file.
4. **Manage Services:**
   - Use commands like `docker-compose down`, `docker-compose logs`, and `docker-compose scale` for managing your application.

## Sample `docker-compose.yml`
```yaml
db:
  image: postgres
  environment:
    POSTGRES_USER: user
    POSTGRES_PASSWORD: password

web:
  image: nginx
  ports:
    - "80:80"
  depends_on:
    - db
```

## Learn More
For detailed information and examples, visit the [Docker Compose Documentation](https://docs.docker.com/compose/).

---

# Terraform

Terraform is an open-source Infrastructure as Code (IaC) tool created by HashiCorp. It allows you to define, provision, and manage cloud infrastructure resources in a safe and efficient manner using declarative configuration files.

## About Terraform
Terraform supports a wide range of cloud providers, including AWS, Azure, Google Cloud, and many others. It enables developers to automate the provisioning and management of infrastructure, ensuring consistency and reducing manual errors.

## Key Features
- **Multi-Cloud Support:** Work seamlessly across multiple cloud providers.
- **Declarative Configuration:** Use HCL (HashiCorp Configuration Language) to define infrastructure.
- **Resource Graphing:** Visualize dependencies between resources.
- **State Management:** Maintain the state of your infrastructure for incremental updates.

## Benefits
- Simplifies infrastructure management by automating provisioning.
- Ensures consistency and repeatability across environments.
- Facilitates collaboration through version-controlled infrastructure definitions.
- Reduces risk by applying changes incrementally with detailed planning.

## How to Use
1. **Install Terraform:**
   - Download and install Terraform from the [official Terraform website](https://www.terraform.io/).
2. **Write Configuration Files:**
   - Create `.tf` files to define resources and their configurations.
3. **Initialize Terraform:**
   - Run `terraform init` to prepare the environment.
4. **Plan and Apply Changes:**
   - Use `terraform plan` to preview changes and `terraform apply` to provision infrastructure.
5. **Manage Infrastructure:**
   - Use `terraform destroy` to clean up resources or `terraform state` commands to manage the state.

## Sample Terraform Configuration
```hcl
provider "aws" {
  region = "us-west-2"
}

resource "aws_instance" "example" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"

  tags = {
    Name = "ExampleInstance"
  }
}
```

## Learn More
For detailed guides and resources, visit the [Terraform Documentation](https://www.terraform.io/docs/).

---

# Docker Commands Cheat Sheet

This README provides a quick reference for commonly used Docker commands, categorized by their purpose. Customize it for your specific needs!

## Network Management
- `docker network ls` - Show all the networks.
- `docker network rm <my_net>` - Delete a specific network.
- `docker create network <name>` - Create a new network.

## Container Management
- `docker ps` - List all running containers.
- `docker ps -a` - List all containers (running and stopped).
- `docker kill` - Kill a running container.
- `docker rm` - Permanently delete a container.
- `docker rm -f` - Forcefully remove a container.
- `docker exec -it <container_id> /bin/bash` - Access a running container's shell.
- `docker logs <container_id>` - View logs from a container.
- `docker stats` - Display real-time statistics of container usage (CPU, memory, etc.).
- `docker stop <container_id>` - Stop a running container.
- `docker start <container_id>` - Start a stopped container.

## Image Management
- `docker images` - Show all Docker images.
- `docker rmi -f` - Forcefully remove a Docker image.
- `docker build -t <image_name> .` - Build a Docker image from a Dockerfile.
- `docker tag <image_id> <repository:tag>` - Tag an image for easier reference.
- `docker pull <repository:tag>` - Download an image from a Docker registry.
- `docker push <repository:tag>` - Upload an image to a Docker registry.

## Volume Management
- `docker volume ls` - Show all volumes.
- `docker volume rm <name_of_volume>` - Delete a specific volume.
- `docker create volume <name>` - Create a new volume.
- `docker volume inspect <name_of_volume>` - View details of a specific volume.

## Additional Commands
- `docker compose up` - Start all services defined in a `docker-compose.yml` file.
- `docker compose down` - Stop all services and remove containers defined in `docker-compose.yml`.
- `docker system prune` - Remove all unused containers, networks, images, and volumes.
- `docker info` - Display detailed information about the Docker installation.
- `docker version` - Show Docker version details.
- `sudo apt install docker-buildx` - Install Docker Buildx.
- `:wq!` - Overwrite and save a file in the vi editor.

---
