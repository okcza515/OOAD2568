import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.HashMap;
import java.util.Map;

public class PayByPayPal implements PaymentStrategy {
    private static final Map<String, String> DATA_BASE = new HashMap<>();
    private final BufferedReader READER = new BufferedReader(new InputStreamReader(System.in));
    private String email;
    private String password;
    private boolean signedIn;

    static {
        DATA_BASE.put("user@example.com", "password123");
        DATA_BASE.put("test@example.com", "testpass");
    }

    @Override
    public void collectPaymentDetails() {
        try {
            while (!signedIn) {
                System.out.print("Enter your PayPal email: ");
                email = READER.readLine();
                System.out.print("Enter your PayPal password: ");
                password = READER.readLine();
                signedIn = authenticate();
            }
        } catch (IOException ex) {
            ex.printStackTrace();
        }
    }

    private boolean authenticate() {
        if (DATA_BASE.containsKey(email) && DATA_BASE.get(email).equals(password)) {
            System.out.println("Authentication successful!");
            return true;
        } else {
            System.out.println("Invalid email or password.");
            return false;
        }
    }

    @Override
    public boolean pay(int paymentAmount) {
        if (signedIn) {
            System.out.println("Paying " + paymentAmount + " using PayPal.");
            return true;
        } else {
            return false;
        }
    }
}
