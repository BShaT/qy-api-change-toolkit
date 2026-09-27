# Paddle setup

Product: **Qy API Change Toolkit**

Payout currency: **CNY**

Price: **¥79 one-time**

Refund window: **7 days**

Support email: `qingy5461@outlook.com`

## 1. Domain Review

Product website:

```text
https://bshat.github.io/qy-api-change-toolkit/
```

The site includes:

```text
Product description
Pricing
Features
Terms
Privacy
Refund policy
Support email
HTTPS
```

In the Paddle dashboard:

```text
Checkout
→ Website approval
→ Submit the GitHub Pages URL
```

Paddle's official guidance says:

- Most submissions are automatically approved.
- Manual review is usually estimated at 5–7 business days.
- Every domain used to launch checkout must be approved.
- The site must be HTTPS and clearly show product, price, features, terms, refund, and privacy.
- Sandbox checkout can be tested before domain approval.

Do not submit an unrelated domain or subdomain.

## 2. Business Identification

Paddle's official guidance says this step is **not required for individuals or sole traders**.

If Paddle asks for a business type, select the individual or sole-trader option if it is available for your account. Do not enter company information that does not exist.

If the account only offers a registered-company path, stop and ask Paddle Support whether an individual seller is supported for your country before submitting documents.

## 3. Identity Verification

For an individual or sole trader, Paddle routes the seller through identity verification in the dashboard or an emailed link.

May be requested through Sumsub:

```text
Government-issued ID
Proof of address
Liveness/selfie check in some cases
```

Complete this only on Paddle or Sumsub's official page. Do not send documents to chat or email the product support address.

Typical manual review is estimated by Paddle at 1–3 business days.

## 4. Tax and payout

In Paddle:

```text
Transfer Preferences
→ Payment Method: Bank/Wire Transfer
→ Transfer Currency: CNY
```

Enter accurate bank details. Paddle currently documents CNY as a supported payout currency.

Paddle's documented payout behavior:

```text
Minimum payout threshold: $100
Monthly payout schedule
Payout starts on the 1st
Funds sent by the 15th
Can take up to 3 working days after sending
Bank or SWIFT fees may apply
```

Keep the payout threshold at the minimum unless you deliberately want a longer accumulation period.

## 5. Product catalog

After domain and identity checks pass, create one product:

```text
Name: Qy API Change Toolkit
Description: OpenAPI, JSON Schema, sitemap, and feed change detection toolkit
Price: ¥79
Billing: One-time
Tax category: Paddle default for digital software
Delivery: Digital download
Refund: 7 days
```

Dashboard path:

```text
Catalog
→ Products
→ New product
```

The API can later create the product and one-time price automatically after a Paddle API key is available.

## 6. API credentials

Create a Paddle API key from the Paddle dashboard after account verification. Store it directly in the Mac mini private configuration file. Never paste it into chat.

The same applies to:

```text
Paddle API key
Webhook secret
Bank details
Identity documents
Tax documents
```

## 7. Sandbox first

Test checkout in Paddle Sandbox before production:

```text
Create sandbox account
→ Create product and ¥79 price
→ Complete a test transaction
→ Verify download delivery
→ Verify webhook
→ Verify 7-day refund rule
```

Sandbox checkout does not require production domain approval.

## 8. Status to send back

```text
Domain Review: pending / approved / rejected
Business Identification: pending / approved / not applicable
Identity Verification: pending / approved / rejected
Tax: pending / completed
Payout CNY: pending / completed
Sandbox test: pending / passed / failed
Production product: pending / created
```

## Official references

- Account verification: https://www.paddle.com/help/start/account-verification
- Domain review: https://www.paddle.com/help/start/account-verification/what-is-domain-verification
- Business identification: https://www.paddle.com/help/start/account-verification/what-is-business-verification
- Identity verification: https://www.paddle.com/help/start/account-verification/what-is-identity-verification
- Payouts: https://www.paddle.com/help/manage/get-paid
- Local-currency payouts: https://www.paddle.com/help/manage/get-paid/can-i-be-paid-in-my-local-currency
