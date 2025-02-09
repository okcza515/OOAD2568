import java.util.Arrays;

public class Main {
    public static void main(String[] args) {
        // Create a customer
        Customer customer = new Customer();
        customer.setName("Napat 65070501074");

        // Create products
        Product product1 = new Product();
        product1.setProductId(1);
        product1.setProductName("Laptop");

        Product product2 = new Product();
        product2.setProductId(2);
        product2.setProductName("Mouse");

        // Create a transaction
        CustomerTransaction transaction = new CustomerTransaction(customer, Arrays.asList(product1, product2));

        // Generate transaction report
        ReportGenerator reportGenerator = new ReportGenerator();
        reportGenerator.generateReport(transaction);
    }
}

//65070501074 Napat Sinjindawong Group 5