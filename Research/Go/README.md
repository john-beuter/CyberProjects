### Purpose

Go Modbus stager allows an attacker to stage multiple attacks. These attacks include a denial of service attack where thousands of Modbus packets overwhelm the human machine interface(HMI) with attacker traffic, preventing the operator from using the HMI, a toggle attack where an attacker can modify the single coil values, and lastly a cycle program where an attacker can cycle the values of a coil between one and zero over a specified time frame. 

### Future Improvements

Store the queued attacks in an array or list that the user can view during run time what attacks are occuring and when they will complete.
For the staged attacks, if the on/off functionality the light remains on or off indefinitely. Add a field for duration: "How long should value remain On/Off? [INSERT TIME IN MIN]" 
