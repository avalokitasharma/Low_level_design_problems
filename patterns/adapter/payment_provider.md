Integrate multiple third-party payment providers (PayTM, Stripe, Razorpay).
Each provider has a different API wrap them to conform to a common Pay(amount) interface.
Your core payment processor should remain agnostic of the provider.
Client requests hold card info, account number, UPI as needed for each payment provider.