
import java.util.ArrayList;
import java.util.List;

public class Main {

    public static void main(String[] args) {
        Customer customer = new Customer("Wavie");
        List<Product> productsList = new ArrayList<>();

        productsList.add(new Product(420, "Crystal Meth Candy"));
        productsList.add(new Product(69, "Mr.White Doritos"));

        CustomerTransaction transaction = new CustomerTransaction(customer, productsList);

        System.out.println(transaction.getName());
        System.out.println(transaction.getDate());
        System.out.println(transaction.productBreakDown());

        AccountsReceivable invoice = new AccountsReceivable(transaction);

        invoice.sendInvoice();
        invoice.postPayment();

        ReportGenerator report = new ReportGenerator(transaction);
        report.generateReport();
    }
}
