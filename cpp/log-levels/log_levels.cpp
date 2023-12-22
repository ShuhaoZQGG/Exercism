#include <string>
using namespace std;
namespace log_line {
    string message(string logMessage){
        string deliminator {" "};
        int index = logMessage.find(deliminator);
        string message = logMessage.substr(index + 1);
        return message;
    }

    string log_level(string logMessage) {
        string deliminator {"]: "};
        int index = logMessage.find(deliminator);
        string level = logMessage.substr(1, index-1);
        return level;
    }

    string reformat(string logMessage) {
        return message(logMessage) + " (" + log_level(logMessage) + ")";
    }
} // namespace log_line