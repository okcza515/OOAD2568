//65070501008
import java.util.ArrayList;
import java.util.List;

public class Main {
    public static void main(String[] args) {
        // Create a customer
        Customer customer = new Customer("Ethan");

        // Create a list of products
        List<Product> products = new ArrayList<>();
        products.add(new Product(101, "Laptop"));
        products.add(new Product(102, "Mouse"));
        products.add(new Product(103, "Keyboard"));

        // Create a transaction
        CustomerTransaction transaction = new CustomerTransaction(customer, products);

        // Process accounts receivable
        AccountsReceivable accountsReceivable = new AccountsReceivable(transaction);
        accountsReceivable.postPayment();
        accountsReceivable.sendInvoice();

        // Generate a transaction report
        ReportGenerator reportGenerator = new ReportGenerator(transaction);
        reportGenerator.generateReport();
    }
}
