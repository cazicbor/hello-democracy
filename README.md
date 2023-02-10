# hello-democracy

Implementation of various voting methods using Golang

## External packages

[cobra](https://pkg.go.dev/github.com/spf13/cobra@v1.5.0) for creating CLI

## Launch project

Run `go run main.go methodCmd` to launch the project

Available CLI (replace `methodCmd` by one of the following):
- majority
- approval
- condorcet
- copeland
- full

## Presentation of the different voting methods

### Two-round system/Majority run-off
Voters cast a single vote for their favorite candidate. If no candidate reveives a simple majority (>50%) during the first round, a second one takes place between the two best-placed candidates

### Approval voting
Takes place in two rounds: each voter selects as many candidates as wanted. The winner is the candidate approved by the largest number of voters. 

### Condorcet method
Selects the candidate who wins a majority of the vote in every head-to-head election against each rival candidate. This method elects a Condorcet winner (if a such winner exists), who is a candidate preferred by more voters than any others. 

### Copeland's method
Each voter ranks candidates in order of preference. A candidate A has majority preference over B if more voters prefer A to B than prefer B to A. For each candidate is calculated the Copeland score which is the number of rivals over whom the candidate is preferred, minus the candidate's number of losses. This method elects a Condorcet winner. 

Article link (in French): https://blog.coddity.com/articles/hello-democracy


