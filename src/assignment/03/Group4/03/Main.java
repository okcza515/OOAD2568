import java.util.ArrayList;
import java.util.List;

public class Main {
    public static void main(String[] args){
        Customer customer = new Customer("Boat");
        List<Product> productsList = new ArrayList<>();
        productsList.add(new Product(1,"Power supply"));
        productsList.add(new Product(2,"HP"));

        CustomerTransaction transaction = new CustomerTransaction(customer,productsList);

        System.out.println( transaction.getName());
        System.out.println( transaction.getDate());
        System.out.println( transaction.productBreakDown());

        AccountsReceivable invoice = new AccountsReceivable(transaction);

        invoice.sendInvoice();
        invoice.postPayment();

        ReportGenerator report = new ReportGenerator(transaction);
        report.generateReport();
    }
}
