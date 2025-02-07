import java.util.ArrayList;
import java.util.List;

public class Main {
    public static void main(String[] args) {
        Customer customer = new Customer("Supakorn Tungpatompramote");

        List<Product> products = new ArrayList<>();
        products.add(new Product(1, "Laptop"));
        products.add(new Product(2, "Smartphone"));

        CustomerTransaction transaction = new CustomerTransaction(customer, products);

        transaction.prepareInvoice();

        PaymentHandler paymentProcessor = new PaymentService();
        AccountsReceivable accountsReceivable = new AccountsReceivable(paymentProcessor);
        accountsReceivable.postPayment();

    
        ReportGenerator report = new ReportGenerator(transaction);
        report.generateReport();
    }
}

//65070501053