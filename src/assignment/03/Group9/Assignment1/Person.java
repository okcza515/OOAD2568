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