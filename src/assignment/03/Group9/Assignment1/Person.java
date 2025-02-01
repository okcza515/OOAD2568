// Sawitt Ngamvilaisiriwong 65070503469
// Single Responsibility Principle: Each interface has one responsibility
interface PowerControl {
    void turnOn();
    void turnOff();
}

interface VolumeControl {
    void volumeUp();
    void volumeDown();
}

// Paratthakon Suksukhon 65070503457
// Interface Segregation Principle: Devices only implement interfaces they need
interface BasicDevice extends PowerControl {
    String getDeviceName();
    boolean isOn();
}
// Paratthakon Suksukhon 65070503457
// Interface for devices that support volume control
interface AudioDevice extends BasicDevice, VolumeControl {
    int getVolume();
}