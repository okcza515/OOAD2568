
public class AccountsReceivable {
	private AccountProcess AccountObject;
	
	public AccountsReceivable(AccountProcess aTransaction){
		AccountObject = aTransaction;
	}
	
	public void postPayment(){
		AccountObject.chargeCustomer();
	}
	
	public void sendInvoice(){
		AccountObject.prepareInvoice();
		// sends the invoice
	}
}
