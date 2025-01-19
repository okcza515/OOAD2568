import java.text.SimpleDateFormat;
import java.util.Date;

class HelloTime {
    public static void main(String[] args) {
        System.out.println("Hello! from 64070501092!");

        while (true) {
            // Get current time
            SimpleDateFormat sdf = new SimpleDateFormat("HH:mm:ss");
            String currentTime = sdf.format(new Date());

            // Print current time
            System.out.println("Current Time: " + currentTime);

            // Wait for 10 seconds
            try {
                Thread.sleep(1000); // 1 seconds in milliseconds
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}
