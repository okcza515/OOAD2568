public class PaymentServiceImplementation implements PaymentService {
    @Override
    public void charge(CustomerTransaction transaction) {
        System.out.println("Charging customer " + transaction.getCustomer().getName());
    }
}
