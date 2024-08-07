### Purpose

Go Modbus stager allows an attacker to stage multiple Modubs attacks to deny legitimate Modbus traffic, send malicous traffic, and interrupt the functionality of PLCs. These attacks include a denial of service attack where thousands of Modbus packets overwhelm the human machine interface(HMI) with attacker traffic, preventing the operator from using the HMI, a On/Off functionality where an attacker can modify the single coil values (turning them from one to zero), and lastly a cycle program where an attacker can cycle the values of a coil between one and zero over a specified time frame. All attacks have the ability to be queued and executed after a set amount of time has passed. With this functionality, attackers can exploit multiple PLCs simultaneously for maximum exploitation of OT infrastructure. 

### Future Improvements

- Store the queued attacks in an array or list that the user can view during run time what attacks are occuring and when they will complete.
- For the staged attacks, if the on/off functionality the light remains on or off indefinitely. Add a field for duration: "How long should value remain On/Off? [INSERT TIME IN MIN]" 
- Make the program use structs instead of variables to store user inputted information
- Potentially allow the attackers to have JSON objects to accept and pars options
