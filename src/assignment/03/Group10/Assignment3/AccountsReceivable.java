// 65070501011 Chayapol Wongpuwarak
public class AccountsReceivable {
	private CustomerTransaction transactionObject;
	private PaymentManager paymentManager;
	
	public AccountsReceivable(CustomerTransaction aTransaction){
		transactionObject = aTransaction;
	}
	
	public void postPayment(){
		paymentManager.chargeCustomer();
	}

	public void sendInvoice(){
		paymentManager.prepareInvoice();
	}
}