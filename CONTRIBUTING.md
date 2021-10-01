# Contributing to Sonar Webhook Operator

:+1::tada: First off, thanks for taking the time to contribute! :tada::+1:

The following is a set of guidelines for contributing to Sonar Webhook Operator and its packages, which are hosted in the [Sonar Webhook Operator Repo](https://github.com/cmwylie19/sonar-webhook-operator) on GitHub. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

#### Table Of Contents

[Code of Conduct](#code-of-conduct)

[Local Development](#local-development)

[Run Project Locally](#run-project-locally)
  * [Atom and Packages](#atom-and-packages)
  * [Atom Design Decisions](#design-decisions)

[How Can I Contribute?](#how-can-i-contribute)
  * [Reporting Bugs](#reporting-bugs)
  * [Suggesting Enhancements](#suggesting-enhancements)
  * [Your First Code Contribution](#your-first-code-contribution)
  * [Pull Requests](#pull-requests)

[Styleguides](#styleguides)
  * [Git Commit Messages](#git-commit-messages)
  * [JavaScript Styleguide](#javascript-styleguide)
  * [CoffeeScript Styleguide](#coffeescript-styleguide)
  * [Specs Styleguide](#specs-styleguide)
  * [Documentation Styleguide](#documentation-styleguide)

[Additional Notes](#additional-notes)
  * [Issue and Pull Request Labels](#issue-and-pull-request-labels)

## Code of Conduct

This project and everyone participating in it is governed by a Code of Coduct. By participating, you are expected to uphold this code. Please report unacceptable behavior to [casewylie@gmail.com](mailto:casewylie@gmail.com).

**Rules**
- Focus on user needs
- Be respectful of other participants
- Strive to have a high level of testing for your code
- Thoroughly test code before submitting a Pull Request

## Local Development
_You will need the following prerequisites for local development._

* Node >= v14
* Golang 1.16.6
* Mongo v5.0.2


## Run Project Locally
```
# clone the project
git clone https://github.com/cmwylie19/sonar-webhook-operator

# Change into the directory
cd sonar-webhook-operator

# Run mongo db through brew or docker
brew services start mongodb-community

# Run the webhook (Go backend)
# In the /webhook directory
go mod tidy
SECRET=secret MONGO_URL=mongodb://localhost:27017 go run main.go    

# Run the frontend
# In the /frontend directory
yarn
REACT_APP_BACKEND_URL=http://localhost:8080 yarn start

```
