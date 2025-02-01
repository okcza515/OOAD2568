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
// Paratthakon Suksukhon 65070503457
// Abstract class to reduce code duplication
abstract class AbstractDevice implements BasicDevice {
    private boolean isOn = false;
    private final String name;

    protected AbstractDevice(String name) {
        this.name = name;
    }

    @Override
    public String getDeviceName() {
        return name;
    }

    @Override
    public boolean isOn() {
        return isOn;
    }

    protected void setOn(boolean on) {
        this.isOn = on;
    }
}