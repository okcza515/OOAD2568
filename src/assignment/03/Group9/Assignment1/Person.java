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

// Warapol Pratumta 65070503466
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

// Warapol Pratumta 65070503466
// Concrete device implementations
class Projector extends AbstractDevice {
    public Projector() {
        super("Projector");
    }

    @Override
    public void turnOn() {
        if (!isOn()) {
            setOn(true);
            System.out.println(getDeviceName() + " is now ON");
        }
    }

    @Override
    public void turnOff() {
        if (isOn()) {
            setOn(false);
            System.out.println(getDeviceName() + " is now OFF");
        }
    }
}

//Chayaphon Chaisangkha 65070503409
// TV implements both power and volume control
class TV extends AbstractDevice implements AudioDevice {
    private int volume = 0;

    public TV() {
        super("TV");
    }

    @Override
    public void turnOn() {
        if (!isOn()) {
            setOn(true);
            System.out.println(getDeviceName() + " is now ON");
        }
    }

    @Override
    public void turnOff() {
        if (isOn()) {
            setOn(false);
            System.out.println(getDeviceName() + " is now OFF");
        }
    }

    @Override
    public void volumeUp() {
        if (isOn() && volume < 100) {
            volume += 5;
            System.out.println(getDeviceName() + " volume: " + volume);
        }
    }

    @Override
    public void volumeDown() {
        if (isOn() && volume > 0) {
            volume -= 5;
            System.out.println(getDeviceName() + " volume: " + volume);
        }
    }

    @Override
    public int getVolume() {
        return volume;
    }
}

class SurroundSoundSystem extends AbstractDevice implements AudioDevice {
    private int volume = 0;

    public SurroundSoundSystem() {
        super("Surround Sound System");
    }

    @Override
    public void turnOn() {
        if (!isOn()) {
            setOn(true);
            System.out.println(getDeviceName() + " is now ON");
        }
    }

    @Override
    public void turnOff() {
        if (isOn()) {
            setOn(false);
            System.out.println(getDeviceName() + " is now OFF");
        }
    }

    @Override
    public void volumeUp() {
        if (isOn() && volume < 100) {
            volume += 5;
            System.out.println(getDeviceName() + " volume: " + volume);
        }
    }

    @Override
    public void volumeDown() {
        if (isOn() && volume > 0) {
            volume -= 5;
            System.out.println(getDeviceName() + " volume: " + volume);
        }
    }

    @Override
    public int getVolume() {
        return volume;
    }
}

//Chanawat Limpanatewin 65070503445
// Open/Closed Principle: RemoteControl is open for extension but closed for modification
class RemoteControl {
    private BasicDevice currentDevice;

    // Dependency Inversion Principle: RemoteControl depends on abstractions
    public void connectTo(BasicDevice device) {
        this.currentDevice = device;
        System.out.println("Remote Control connected to: " + device.getDeviceName());
    }

    public void powerOn() {
        if (currentDevice != null) {
            currentDevice.turnOn();
        } else {
            System.out.println("No device connected");
        }
    }

    public void powerOff() {
        if (currentDevice != null) {
            currentDevice.turnOff();
        } else {
            System.out.println("No device connected");
        }
    }

    public void volumeUp() {
        if (currentDevice instanceof AudioDevice) {
            ((AudioDevice) currentDevice).volumeUp();
        } else {
            System.out.println("Current device does not support volume control");
        }
    }

    public void volumeDown() {
        if (currentDevice instanceof AudioDevice) {
            ((AudioDevice) currentDevice).volumeDown();
        } else {
            System.out.println("Current device does not support volume control");
        }
    }

    public BasicDevice getCurrentDevice() {
        return currentDevice;
    }
}

//Chanawat Limpanatewin 65070503445
// Person class to test the system
class Person {
    public static void main(String[] args) {
        // Initialize devices
        TV tv = new TV();
        Projector projector = new Projector();
        SurroundSoundSystem soundSystem = new SurroundSoundSystem();

        // Create remote control
        RemoteControl remote = new RemoteControl();

        // Test TV (supports both power and volume)
        System.out.println("\nTesting TV:");
        remote.connectTo(tv);
        remote.powerOn();
        remote.volumeUp();
        remote.volumeUp();
        remote.volumeDown();
        remote.powerOff();

        // Test Projector (supports only power)
        System.out.println("\nTesting Projector:");
        remote.connectTo(projector);
        remote.powerOn();
        remote.volumeUp(); // Should show error message
        remote.powerOff();

        // Test Sound System (supports both power and volume)
        System.out.println("\nTesting Sound System:");
        remote.connectTo(soundSystem);
        remote.powerOn();
        remote.volumeUp();
        remote.volumeUp();
        remote.volumeDown();
        remote.powerOff();
    }
}