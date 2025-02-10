public class CustomerTransactionReportGenerator implements ReportGenerator {
    @Override
    public void generateReport(CustomerTransaction transaction) {
        StringBuilder report = new StringBuilder();
        report.append("Transaction Report for ")
              .append(transaction.getCustomer().getName())
              .append("\nDate: ")
              .append(transaction.getDate())
              .append("\nProducts:\n");
        for (Product p : transaction.getProducts()) {
            report.append("- ")
                  .append(p.getProductName())
                  .append(" (ID: ")
                  .append(p.getProductId())
                  .append(")\n");
        }
        System.out.println(report.toString());
    }
}
