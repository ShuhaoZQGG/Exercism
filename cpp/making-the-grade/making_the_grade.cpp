#include <array>
#include <string>
#include <vector>
#include <cmath>
// Round down all provided student scores.
std::vector<int> round_down_scores(std::vector<double> student_scores) {
    // TODO: Implement round_down_scores
    std::vector<int> result {};
    for (double score : student_scores) {
      int rounded_score = static_cast<int>(score);
      result.push_back(rounded_score);
    }
    return result;
}


// Count the number of failing students out of the group provided.
int count_failed_students(std::vector<int> student_scores) {
    // TODO: Implement count_failed_students
    int cut_off = 40;
    int count = 0;
    for (int score : student_scores) {
      if (score <= cut_off) {
        count ++;
      }
    }
    return count;
}

// Determine how many of the provided student scores were 'the best' based on the provided threshold.
std::vector<int> above_threshold(std::vector<int> student_scores, int threshold) {
    // TODO: Implement above_threshold
    std::vector<int> result {};
    for (int score : student_scores) {
      if (score >= threshold) {
        result.push_back(score);
      }
    }
    return result;
}

// Create a list of grade thresholds based on the provided highest grade.
std::array<int, 4> letter_grades(int highest_score) {
    // TODO: Implement letter_grades
    double lowest_score { 40 };
    double difference = highest_score - lowest_score;
    double difference_slice = difference / 4;
    int difference_slice_int {static_cast<int>(std::floor(difference_slice))};
    int highest_bracket {};
    if (difference_slice != std::floor(difference_slice)) {
      highest_bracket = highest_score - difference_slice_int;
    } else {
      highest_bracket = highest_score - difference_slice + 1;
    }
    std::array<int, 4> brackets {};
    for (int i = 0; i < 4; i++) {
      brackets.at(i) = highest_bracket - (3 - i) * difference_slice_int;
    }
    return brackets;
}

// Organize the student's rank, name, and grade information in ascending order.
std::vector<std::string> student_ranking(std::vector<int> student_scores, std::vector<std::string> student_names) {
    // TODO: Implement student_ranking
    std::vector<std::string> result {};
    for (int i = 0; i < student_scores.size(); i++) {
      std::string name = student_names[i];
      int score = student_scores[i];
      std::string content = { std::to_string(i+1) + ". " + name + ": " + std::to_string(score)};
      result.push_back(content);
    }

    return result;
}

// Create a string that contains the name of the first student to make a perfect score on the exam.
std::string perfect_score(std::vector<int> student_scores, std::vector<std::string> student_names) {
    // TODO: Implement perfect_score
    for (int i = 0; i < student_scores.size(); ++i) {
      if (student_scores[i] == 100) {
        return student_names[i];
      }
    }
    return "";
}
