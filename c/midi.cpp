#include <iostream>    
#include <CoreMIDI/CoreMIDI.h>
int main() {} 
// Initialize the Core MIDI framework
      MIDIClientRef midiClient;
      MIDIClientCreate(CFSTR("MIDIDeviceList"), NULL, NULL, &midiClient);
      // Get the number of MIDI devices                                                                                                                                                       
      ItemCount numDevices = MIDIGetNumberOfDevices();
      if (numDevices == 0) {
	      std::cout << "No MIDI devices found." << std::endl;
	      return 0;
      }
      std::cout << "MIDI Devices:" << std::endl;
      for (ItemCount i = 0; i < numDevices; ++i) {
	      MIDIDeviceRef midiDevice = MIDIGetDevice(i);
	      // Get the name of the MIDI device

	      CFStringRef deviceName;
	      MIDIObjectGetStringProperty(midiDevice, kMIDIPropertyName, &deviceName);
	      std::cout << "Device " << i << ": " << CFStringGetCStringPtr(deviceName, kCFStringEncodingUTF8) << std::endl;   
      }
	// Cleanup

	MIDIClientDispose(midiClient);

	return 0;                                                                                                  
}                  
