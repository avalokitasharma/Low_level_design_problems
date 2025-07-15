Implement a ride fare calculator that can switch between pricing strategies: regular, night, and surge.
 - Each strategy may require different configuration parameters:
 - SurgeFare might need a struct with city-specific factors, time slots, etc.
 - NightFare might just take an integer multiplier.

We want to avoid coupling FareCalculator to strategy internals.