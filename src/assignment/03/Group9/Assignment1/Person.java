// Single Responsibility Principle: Each interface has one responsibility
interface PowerControl {
    void turnOn();
    void turnOff();
}

interface VolumeControl {
    void volumeUp();
    void volumeDown();
}
