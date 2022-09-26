package electionday

import "strconv"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	pInitialVotes := new(int)
	*pInitialVotes = initialVotes
	return pInitialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}

	return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	if len(candidateName) <= 0 {
		panic("candidate name is invalid")
	}

	electionResult := new(ElectionResult)
	electionResult.Name = candidateName
	electionResult.Votes = votes

	return electionResult
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	if result == nil {
		panic("input is nil")
	}

	message := result.Name + " (" + strconv.Itoa(result.Votes) + ")"
	return message
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {

	if results == nil {
		panic("results are nil")
	}

	if len(candidate) <= 0 {
		panic("candidate name is invalid")
	}

	_, found := results[candidate]
	if !found {
		panic("no such candidate in map")
	}

	results[candidate]--
}
