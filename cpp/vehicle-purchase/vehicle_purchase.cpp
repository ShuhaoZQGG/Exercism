#include "vehicle_purchase.h"
#include <string>
using namespace std;
namespace vehicle_purchase {

    // needs_license determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
    bool needs_license(string kind){
        // TODO: Return true if you need a license for that kind of vehicle.
        if (kind == "car" or kind == "truck") {
            return true;
        }
        return false;
    }

    // choose_vehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
    string choose_vehicle(string option1, string option2) {
        // TODO: Return the final decision in a sentence.
        string postfix = {" is clearly the better choice."};
        if (option1 < option2) {
            return option1 + postfix;
        }
        return option2 + postfix;
    }

    // calculate_resell_price calculates how much a vehicle can resell for at a certain age.
    double calculate_resell_price(double original_price, double age) {
        // TODO: Return the age-corrected resell price.
        double depreciation {0.8};
        double price {};
        if (age >= 3 and age < 10) {
            depreciation = 0.7;
            price = depreciation * original_price;
        } else if (age >= 10) {
            depreciation = 0.5;
            price = depreciation * original_price;
        } else {
            price = depreciation * original_price;
        }

        return price;
    }

}  // namespace vehicle_purchase