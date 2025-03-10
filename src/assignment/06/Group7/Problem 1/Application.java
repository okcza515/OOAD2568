import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.HashMap;
import java.util.Map;

public class Application {
    private static final Map<Integer, Integer> priceOnProducts = new HashMap<>();

    static {
        priceOnProducts.put(1, 2200); // Motherboard
        priceOnProducts.put(2, 1850); // CPU
        priceOnProducts.put(3, 1100); // HDD
        priceOnProducts.put(4, 890);  // Memory
    }

    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        Order order = new Order();
        PaymentStrategy strategy = null;

        while (!order.isClosed()) {
            int cost;
            String continueChoice;
            
            do {
                System.out.print("Please, select a product:\n" +
                        "1 - Motherboard (2200 units)\n" +
                        "2 - CPU (1850 units)\n" +
                        "3 - HDD (1100 units)\n" +
                        "4 - Memory (890 units)\n");
                int choice = Integer.parseInt(reader.readLine());
                cost = priceOnProducts.get(choice);
                System.out.print("Count: ");
                int count = Integer.parseInt(reader.readLine());
                order.setTotalCost(cost * count);
                System.out.print("Do you wish to continue selecting products? (Y/N): ");
                continueChoice = reader.readLine();
            } while (continueChoice.equalsIgnoreCase("Y"));

            // Select payment method
            System.out.println("Choose payment method: 1 - PayPal, 2 - Credit Card");
            String paymentChoice = reader.readLine();

            if (paymentChoice.equals("1")) {
                strategy = new PayByPayPal();
            } else if (paymentChoice.equals("2")) {
                strategy = new PayByCreditCard();
            } else {
                System.out.println("Invalid payment method.");
                continue;
            }

            order.processOrder(strategy);

            System.out.print("Pay " + order.getTotalCost() + " units or Continue shopping? (P/C): ");
            String proceed = reader.readLine();
            if (proceed.equalsIgnoreCase("P")) {
                if (order.pay(strategy)) {
                    System.out.println("Payment has been successful.");
                    order.setClosed();
                } else {
                    System.out.println("Payment failed. Please check your details.");
                }
            }
        }
    }
}