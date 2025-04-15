
// Creates the magic packet bytes given a mac address.
char *makeMagic(char *mac);

// Sends the created magic packet bytes to the broadcast address.
void sendMagic(char *magic);

// Prints a mac address in hex
void printMacAddr(char *mac);

// Given a string, validate it has a mac format and return an integer with the
// hex values
int *parseMacAddr(char *mac);
