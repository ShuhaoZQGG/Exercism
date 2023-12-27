#include <string>
#include <vector>
#include <iostream>
namespace election {

// The election result struct is already created for you:

struct ElectionResult {
    // Name of the candidate
    std::string name{};
    // Number of votes the candidate has
    int votes{};
};

// TODO: Task 1
// vote_count takes a reference to an `ElectionResult` as an argument and will
// return the number of votes in the `ElectionResult.
int vote_count (ElectionResult& election) {
  return election.votes;
}

// TODO: Task 2
// increment_vote_count takes a reference to an `ElectionResult` as an argument
// and a number of votes (int), and will increment the `ElectionResult` by that
// number of votes.
void increment_vote_count(ElectionResult& election, int new_votes) {
  election.votes += new_votes;
}


// TODO: Task 3
// determine_result receives the reference to a final_count and returns a
// reference to the `ElectionResult` of the new president. It also changes the
// name of the winner by prefixing it with "President". The final count is given
// in the form of a `reference` to `std::vector<ElectionResult>`, a vector with
// `ElectionResults` of all the participating candidates.
ElectionResult& determine_result(std::vector<ElectionResult>& electionResultVector) {
    // Check for an empty vector and throw an exception if necessary.
    if (electionResultVector.empty()) {
        throw std::runtime_error("The election result vector is empty.");
    }

    // Initialize variables to track the index of the candidate with the most votes.
    int presidentIndex = 0;

    // Iterate over the candidates to find the index of the one with the most votes.
    for (size_t i = 0; i < electionResultVector.size(); ++i) {
        if (electionResultVector[i].votes > electionResultVector[presidentIndex].votes) {
            presidentIndex = i;
        }
    }

    // Append "President" to the name of the candidate with the most votes.
    electionResultVector[presidentIndex].name = "President " + electionResultVector[presidentIndex].name;

    // Return a reference to the winning candidate.
    return electionResultVector[presidentIndex];
}

}  // namespace election