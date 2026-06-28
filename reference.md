# Reference
## Clients
<details><summary><code>client.Clients.ListClients() -> *v2.PaginatedResponseClientResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of all clients associated with your partner account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListClientsRequest{}
client.Clients.ListClients(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` — Field to order the results by, e.g., 'created_at:desc,updated_at:asc'
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: id, name, correlation_id, company_number, status. Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.CreateClient(request) -> *v2.ClientResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new client under your partner account. The client will remain in a pending state until reviewed by Winyield.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ClientCreatePayload{
        Name: "ACME Corp",
        Jurisdiction: v2.JurisdictionEnumEu,
    }
client.Clients.CreateClient(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**correlationID:** `*string` — The correlation ID you provided at the creation of the client
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the client
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*v2.ClientTypeEnum` — The type of the client, must be one of `individual`,`corporate`. Default is `corporate` if not provided.
    
</dd>
</dl>

<dl>
<dd>

**jurisdiction:** `*v2.JurisdictionEnum` — The jurisdiction of the client, must be one of `eu`, `us`, `uk`
    
</dd>
</dl>

<dl>
<dd>

**companyNumber:** `*string` — The company number of the client if type is `corporate`
    
</dd>
</dl>

<dl>
<dd>

**data:** `map[string]any` — Additional data associated with the client
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.CreateClientData(request) -> *v2.ClientDataResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Upload supplementary client information, such as bank account balance, accounting figures, or other relevant details.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ClientDataCreatePayload{
        ClientID: "client_123",
        Data: map[string]any{
            "key1": "value1",
            "key2": "value2",
        },
    }
client.Clients.CreateClientData(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**data:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.ListLimitRequests() -> *v2.PaginatedResponseLimitRequestResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of all limit requests associated with your partner account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListLimitRequestsRequest{}
client.Clients.ListLimitRequests(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `*string` — Filter by client ID
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` — Field to order the results by, e.g., 'created_at:desc,updated_at:asc'
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: id, client_id. Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.CreateLimitRequest(request) -> *v2.LimitRequestResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a limit review request for a client. The request will remain in a pending state until reviewed by Winyield.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.LimitRequestCreatePayload{
        ClientID: "client_1234567890abcdef",
        RequestedLimit: &v2.LimitRequestCreatePayloadRequestedLimit{
            Double: 1.1,
        },
        Reason: "Need more credit for business expansion",
    }
client.Clients.CreateLimitRequest(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` — The ID of the client for which the limit request is being created
    
</dd>
</dl>

<dl>
<dd>

**requestedLimit:** `*v2.LimitRequestCreatePayloadRequestedLimit` — The requested credit limit amount
    
</dd>
</dl>

<dl>
<dd>

**reason:** `string` — The reason for the limit request
    
</dd>
</dl>

<dl>
<dd>

**waiverRequest:** `*bool` — Whether a special waiver is requested alongside this limit request
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.GetLimitRequest(RequestID) -> *v2.LimitRequestResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a specific limit request by its ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.GetLimitRequestRequest{
        RequestID: "request_id",
    }
client.Clients.GetLimitRequest(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**requestID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.ListOnboardingClients() -> *v2.PaginatedResponseClientResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all clients that have self-registered via the portal and are awaiting partner approval.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListOnboardingClientsRequest{}
client.Clients.ListOnboardingClients(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.ApproveOnboarding(ClientID) -> *v2.ClientResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Accept a self-onboarded client. The client status moves to 'pending' and the owner's portal account is activated so they can begin submitting documents.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ApproveOnboardingRequest{
        ClientID: "client_id",
    }
client.Clients.ApproveOnboarding(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.RejectOnboarding(ClientID) -> *v2.ClientResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Reject a self-onboarded client. The client record is kept with 'rejected' status for audit history and is not deleted.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.RejectOnboardingRequest{
        ClientID: "client_id",
    }
client.Clients.RejectOnboarding(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.AddClientPortalUser(ClientID, request) -> *v2.ClientUserResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Invite a new user to a client's portal account. The invited user will receive an email with a one-time link to set their password. Partner can assign any role: 'owner', 'admin', or 'viewer'.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ClientUserInviteRequest{
        ClientID: "client_id",
        FirstName: "first_name",
        LastName: "last_name",
        Email: "email",
        RoleType: "role_type",
    }
client.Clients.AddClientPortalUser(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**firstName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**lastName:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**phone:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**roleType:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.ListClientWaivers(ClientID) -> *v2.PaginatedResponseWaiverResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all waivers associated with a specific client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListClientWaiversRequest{
        ClientID: "client_id",
    }
client.Clients.ListClientWaivers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` — Field to order the results by, e.g., 'created_at:desc'
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: id, client_id, status. Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.GetClientByID(ClientID) -> *v2.ClientResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve detailed information for a specific client using their client ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.GetClientByIDRequest{
        ClientID: "client_id",
    }
client.Clients.GetClientByID(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.DeleteClient(ClientID) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a client by ID. Only clients with 'pending' status can be deleted. All client's loans must also be in 'pending' status.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.DeleteClientRequest{
        ClientID: "client_id",
    }
client.Clients.DeleteClient(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.ListClientChecklistSummaries(ClientID) -> *v2.PaginatedResponseChecklistSummaryPartnerResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the checklist summaries for one of your clients, including the AI description and item completion counts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListClientChecklistSummariesRequest{
        ClientID: "client_id",
    }
client.Clients.ListClientChecklistSummaries(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sandbox
<details><summary><code>client.Sandbox.UpdateClient(ClientID, request) -> *v2.ClientResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update an existing client's status or credit limit using their client ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ClientUpdateSandbox{
        ClientID: "client_id",
    }
client.Sandbox.UpdateClient(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*v2.ClientStatusEnum` — The status of the client. One of the following: `active, rejected, deactivated, pending, pending_onboarding, pre_approved, deleted, inactive`
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*v2.ClientUpdateSandboxLimit` — The limit to set for the client. This will override the existing limit.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandbox.UpdateLoan(LoanID, request) -> *v2.LoanResponseWithInstallments</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the status of a specific loan using its unique loan ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.LoanUpdateSandbox{
        LoanID: "loan_id",
    }
client.Sandbox.UpdateLoan(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**loanID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*v2.LoanStatusEnum` — The status of the client. One of the following: `pending, overdue, active, default, sold, restructured, repaid, pre_approved, rejected, deleted, inactive`
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sandbox.WebhookTest(request) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Test a webhook subscription by ID or trigger all by event type.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.WebhookTestSandbox{
        EventType: v2.WebhookEventTypeEnumLoanUpdated,
    }
client.Sandbox.WebhookTest(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookID:** `*string` — The ID of the webhook subscription. Only this webhook will be triggered if provided.
    
</dd>
</dl>

<dl>
<dd>

**eventType:** `*v2.WebhookEventTypeEnum` — Event type to trigger for the test. All subscriptions for this event type will be triggered if webhook_id is not provided.Possible values: loan_updated, installment_updated, client_updated, client_limit_updated, partner_limit_updated
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounts
<details><summary><code>client.Accounts.ListClientAccountFields(ClientID) -> map[string]*v2.CurrencyFieldSpec</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return the required and optional bank account fields for each supported currency. Fetch once and cache; use it to build the create-account request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListClientAccountFieldsRequest{
        ClientID: "client_id",
    }
client.Accounts.ListClientAccountFields(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.ListClientAccounts(ClientID) -> *v2.PaginatedResponseClientAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all bank accounts for one of your clients.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListClientAccountsRequest{
        ClientID: "client_id",
    }
client.Accounts.ListClientAccounts(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.CreateClientAccount(ClientID, request) -> *v2.ClientAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a bank account for one of your clients. The account is registered with the payment provider immediately. Use the `status` field to create it as `active` (default; demotes any existing active account in the same currency to `passive`) or `passive`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.PartnerClientAccountCreateRequest{
        ClientID: "client_id",
        AccountHolderName: "Acme Ltd",
        AccountHolderType: v2.AccountHolderTypeEnumBusiness,
        Currency: v2.CurrencyEnumGbp,
        SortCode: v2.String(
            "40-47-84",
        ),
        AccountNumber: v2.String(
            "12345678",
        ),
    }
client.Accounts.CreateClientAccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**accountHolderName:** `string` — Full name of the account holder.
    
</dd>
</dl>

<dl>
<dd>

**label:** `*string` — Optional label / nickname for the account (max 50 characters).
    
</dd>
</dl>

<dl>
<dd>

**accountHolderType:** `*v2.AccountHolderTypeEnum` — Account holder type. One of: `business`, `private`.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*v2.CurrencyEnum` — ISO 4217 currency code. Use `/accounts/fields` to get required fields per currency.
    
</dd>
</dl>

<dl>
<dd>

**sortCode:** `*string` — Sort code (required for GBP).
    
</dd>
</dl>

<dl>
<dd>

**accountNumber:** `*string` — Account number (required for GBP and USD).
    
</dd>
</dl>

<dl>
<dd>

**iban:** `*string` — IBAN (required for EUR, CZK, PLN).
    
</dd>
</dl>

<dl>
<dd>

**bic:** `*string` — BIC / SWIFT code (optional for EUR).
    
</dd>
</dl>

<dl>
<dd>

**routingNumber:** `*string` — ABA routing number (required for USD).
    
</dd>
</dl>

<dl>
<dd>

**accountType:** `*string` — Account type (required for USD). E.g. `checking` or `savings`.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*v2.AccountAddress` — Account holder address (required for USD).
    
</dd>
</dl>

<dl>
<dd>

**status:** `*v2.PartnerClientAccountCreateRequestStatus` — Account status. `active` demotes any existing active account in the same currency to `passive`; `passive` is added as a backup. Defaults to `active`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounts.GetClientAccount(ClientID, AccountID) -> *v2.ClientAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a specific bank account for one of your clients.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.GetClientAccountRequest{
        ClientID: "client_id",
        AccountID: "account_id",
    }
client.Accounts.GetClientAccount(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**accountID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Documents
<details><summary><code>client.Documents.ListDocuments() -> *v2.PaginatedResponseDocumentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all documents linked to a client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListDocumentsRequest{}
client.Documents.ListDocuments(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**loanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**installmentID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**waterfallID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` — Field to order the results by, e.g., 'created_at:desc,updated_at:asc'
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: id, client_id, loan_id, installment_id, waterfall_id, category, file_name, document_date, folder_path. Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Documents.UploadDocument(request) -> *v2.DocumentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Upload a new document related to a client or loan, such as financial statements or KYC files.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.DocumentCreatePayload{
        File: strings.NewReader(
            "",
        ),
        Category: "category",
        FileName: "file_name",
    }
client.Documents.UploadDocument(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**loanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**installmentID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**waterfallID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Documents.GetAvailableDocumentCategories() -> *v2.AvailableDocumentCategoriesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all available document categories.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Documents.GetAvailableDocumentCategories(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Documents.GetDocumentByID(DocumentID) -> *v2.DocumentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details for a specific document using its document ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.GetDocumentByIDRequest{
        DocumentID: "document_id",
    }
client.Documents.GetDocumentByID(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**documentID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Documents.DeleteDocument(DocumentID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific document by using its document ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.DeleteDocumentRequest{
        DocumentID: "document_id",
    }
client.Documents.DeleteDocument(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**documentID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Investors
<details><summary><code>client.Investors.InvestorListClients() -> *v2.PaginatedResponseClientInvestorResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all clients with at least one loan funded by this investor.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.InvestorListClientsRequest{}
client.Investors.InvestorListClients(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: id, name, correlation_id, company_number, status. Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Investors.InvestorGetClient(ClientID) -> *v2.ClientInvestorResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a specific client that has a loan funded by this investor.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.InvestorGetClientRequest{
        ClientID: "client_id",
    }
client.Investors.InvestorGetClient(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Investors.InvestorListLoans() -> *v2.PaginatedResponseLoanInvestorResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all loans funded by the current investor.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.InvestorListLoansRequest{}
client.Investors.InvestorListLoans(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: id, partner_id, client_id, status, loan_date, currency, partner.name, client.name. Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Investors.InvestorGetLoan(LoanID) -> *v2.LoanResponseWithInstallments</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a specific loan funded by the current investor, with its installments.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.InvestorGetLoanRequest{
        LoanID: "loan_id",
    }
client.Investors.InvestorGetLoan(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**loanID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Investors.InvestorListInstallments() -> *v2.PaginatedResponseInstallmentResponseWithClientInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all installments for loans funded by the current investor.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.InvestorListInstallmentsRequest{}
client.Investors.InvestorListInstallments(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**loanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: id, client_id, loan_id, status, client.name, expected_repayment_at. Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Investors.InvestorGetInstallment(InstallmentID) -> *v2.InstallmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a specific installment for a loan funded by the current investor.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.InvestorGetInstallmentRequest{
        InstallmentID: "installment_id",
    }
client.Investors.InvestorGetInstallment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**installmentID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Investors.InvestorListRepayments() -> *v2.PaginatedResponseRepaymentResponseWithClientInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all repayment logs for loans funded by the current investor.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.InvestorListRepaymentsRequest{}
client.Investors.InvestorListRepayments(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**loanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**installmentID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: id, client_id, loan_id, installment_id, created_at, client.name, client.correlation_id. Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Installments
<details><summary><code>client.Installments.ListInstallments() -> *v2.PaginatedResponseInstallmentResponseWithClientInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of all loan installments under your partner account. Supports optional filtering by loan or client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListInstallmentsRequest{}
client.Installments.ListInstallments(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**loanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` — Field to order the results by, e.g., 'created_at:desc,updated_at:asc'
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: id, client_id, loan_id, status, client.name, expected_repayment_at. Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Installments.AddInstallment(request) -> []*v2.InstallmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add new installments to a loan with its specific loan ID. This endpoint is available to select partners and will trigger the recalculation of the IRR and interest amounts for all installments of the loan.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.InstallmentCreatePayload{
        LoanID: "loan_12345",
        Installments: []*v2.LoanInstallmentCreatePayload{
            &v2.LoanInstallmentCreatePayload{
                ExpectedRepaymentAt: v2.MustParseDate(
                    "2025-12-01",
                ),
                Amount: &v2.LoanInstallmentCreatePayloadAmount{
                    String: "1000.00",
                },
            },
            &v2.LoanInstallmentCreatePayload{
                ExpectedRepaymentAt: v2.MustParseDate(
                    "2026-01-01",
                ),
                Amount: &v2.LoanInstallmentCreatePayloadAmount{
                    String: "1000.00",
                },
            },
        },
    }
client.Installments.AddInstallment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**loanID:** `string` — The loan ID to add the installments to
    
</dd>
</dl>

<dl>
<dd>

**installments:** `[]*v2.LoanInstallmentCreatePayload` — List of installments to add to the loan
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Installments.ListPaymentPromises() -> *v2.PaginatedResponsePaymentPromiseResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of payment promises recorded for installments under your partner account. Supports optional filtering by loan or client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListPaymentPromisesRequest{}
client.Installments.ListPaymentPromises(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` — Field to order the results by, e.g., 'created_at:desc,updated_at:asc'
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: id, installment_id, status, promised_date, created_at. Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>

<dl>
<dd>

**loanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Installments.CreatePaymentPromise(request) -> *v2.PaymentPromiseResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Record a payment promise made by a client for one of their installments. The promised date must be today or in the future.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.PaymentPromiseCreatePayload{
        InstallmentID: "inst_12345",
        Amount: &v2.PaymentPromiseCreatePayloadAmount{
            Double: 1.1,
        },
        PromisedDate: v2.MustParseDate(
            "2026-05-15",
        ),
    }
client.Installments.CreatePaymentPromise(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**installmentID:** `string` — The ID of the installment this promise relates to
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*v2.PaymentPromiseCreatePayloadAmount` — The amount the client has promised to pay (must be > 0)
    
</dd>
</dl>

<dl>
<dd>

**promisedDate:** `time.Time` — The date the client has committed to pay by (today or future)
    
</dd>
</dl>

<dl>
<dd>

**notes:** `*string` — Optional notes captured during the collections interaction
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Installments.GetInstallmentByID(InstallmentID) -> *v2.InstallmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve detailed information for a specific installment using its installment ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.GetInstallmentByIDRequest{
        InstallmentID: "installment_id",
    }
client.Installments.GetInstallmentByID(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**installmentID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Installments.EditInstallment(InstallmentID, request) -> *v2.InstallmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update an installment's amount and expected repayment date with its specific installment ID. This endpoint is available to select partners and will trigger the recalculation of the IRR and interest amounts for all installments of the loan.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.InstallmentEditPayload{
        InstallmentID: "installment_id",
        Amount: &v2.InstallmentEditPayloadAmount{
            Double: 1.1,
        },
        ExpectedRepaymentAt: v2.MustParseDate(
            "2025-12-01",
        ),
    }
client.Installments.EditInstallment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**installmentID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*v2.InstallmentEditPayloadAmount` — The new amount for the installment
    
</dd>
</dl>

<dl>
<dd>

**expectedRepaymentAt:** `time.Time` — The new expected repayment date
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Installments.DeleteInstallment(InstallmentID) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete an installment with its specific installment ID. This endpoint is available to select partners and will trigger the recalculation of the IRR and interest amounts for all installments of the loan.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.DeleteInstallmentRequest{
        InstallmentID: "installment_id",
    }
client.Installments.DeleteInstallment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**installmentID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Loans
<details><summary><code>client.Loans.ListLoans() -> *v2.PaginatedResponseLoanResponseWithClientInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all loans associated with your partner account. Supports optional filtering by client ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListLoansRequest{}
client.Loans.ListLoans(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` — Field to order the results by, e.g., 'created_at:desc,updated_at:asc'
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: id, client_id, status, client.name, correlation_id. Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loans.CreateLoan(request) -> *v2.LoanResponseWithInstallments</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new loan for an approved client with an active credit limit.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.LoanCreatePayload{
        ClientID: "client_ACME",
        Currency: v2.CurrencyEnumEur,
        Amount: &v2.LoanCreatePayloadAmount{
            Double: 1.1,
        },
        Installments: []*v2.LoanInstallmentCreatePayload{
            &v2.LoanInstallmentCreatePayload{
                ExpectedRepaymentAt: v2.MustParseDate(
                    "2025-12-01",
                ),
                Amount: &v2.LoanInstallmentCreatePayloadAmount{
                    Double: 1.1,
                },
            },
        },
    }
client.Loans.CreateLoan(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` — The ID of the client for this loan
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*v2.CurrencyEnum` — The currency of the loan, must be one of the supported currencies: eur, gbp, usd, czk, pln, isk
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*v2.LoanCreatePayloadAmount` — The amount of the loan
    
</dd>
</dl>

<dl>
<dd>

**correlationID:** `*string` — The correlation ID you provided at the creation of the loan
    
</dd>
</dl>

<dl>
<dd>

**loanDate:** `*time.Time` — Please provide the loan_date if it differs from the loan creation time (created_at). Otherwise, it will be automatically set.
    
</dd>
</dl>

<dl>
<dd>

**installments:** `[]*v2.LoanInstallmentCreatePayload` — List of installments for the loan, each with a due date and amount
    
</dd>
</dl>

<dl>
<dd>

**data:** `map[string]any` — Additional data related to the loan
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loans.GetLoanByID(LoanID) -> *v2.LoanResponseWithInstallments</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve detailed information about a specific loan by its loan ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.GetLoanByIDRequest{
        LoanID: "loan_id",
    }
client.Loans.GetLoanByID(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**loanID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loans.DeleteLoan(LoanID) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a loan by ID. Only loans with 'pending' status can be deleted.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.DeleteLoanRequest{
        LoanID: "loan_id",
    }
client.Loans.DeleteLoan(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**loanID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loans.CreateBulkLoans(request) -> *v2.BulkLoanTaskResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create multiple loans in a single request. Processing happens asynchronously. Returns a task ID for tracking progress.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.BulkLoanCreatePayload{
        Loans: []*v2.BulkLoanItemPayload{
            &v2.BulkLoanItemPayload{
                ClientID: "client_123",
                Currency: v2.CurrencyEnumEur,
                Amount: &v2.BulkLoanItemPayloadAmount{
                    String: "50000.00",
                },
                CorrelationID: v2.String(
                    "LOAN_001",
                ),
                LoanDate: v2.Time(
                    v2.MustParseDate(
                        "2023-05-01",
                    ),
                ),
                Installments: []*v2.LoanInstallmentCreatePayload{
                    &v2.LoanInstallmentCreatePayload{
                        ExpectedRepaymentAt: v2.MustParseDate(
                            "2023-06-01",
                        ),
                        Amount: &v2.LoanInstallmentCreatePayloadAmount{
                            String: "26000.00",
                        },
                    },
                    &v2.LoanInstallmentCreatePayload{
                        ExpectedRepaymentAt: v2.MustParseDate(
                            "2023-07-01",
                        ),
                        Amount: &v2.LoanInstallmentCreatePayloadAmount{
                            String: "26000.00",
                        },
                    },
                },
                Data: map[string]any{
                    "loan_type": "business",
                    "purpose": "expansion",
                },
            },
        },
    }
client.Loans.CreateBulkLoans(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**loans:** `[]*v2.BulkLoanItemPayload` — List of loans to create (max 1000)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loans.GetBulkLoanStatus(TaskID) -> *v2.BulkLoanTaskStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Check the status of a bulk loan creation task and retrieve results when completed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.GetBulkLoanStatusRequest{
        TaskID: "task_id",
    }
client.Loans.GetBulkLoanStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**taskID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Loans.SetLoanDefault(LoanID, request) -> *v2.LoanResponseWithInstallments</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Mark a loan as defaulted, recording the default date and the amount recovered from selling it. Defaults the loan's active and overdue installments and updates the loan status accordingly.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.LoanDefaultPayload{
        LoanID: "loan_id",
        DefaultDate: v2.MustParseDate(
            "2026-06-23",
        ),
        SoldAmount: &v2.LoanDefaultPayloadSoldAmount{
            Double: 1.1,
        },
    }
client.Loans.SetLoanDefault(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**loanID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**defaultDate:** `time.Time` — Date the loan is marked as defaulted.
    
</dd>
</dl>

<dl>
<dd>

**soldAmount:** `*v2.LoanDefaultPayloadSoldAmount` — Amount recovered when the defaulted loan is sold.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Partners
<details><summary><code>client.Partners.GetAvailableFunding() -> []*v2.AvailableFunding</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this endpoint to check the available funding capacity in your dedicated lending account before initiating a new loan or submitting a drawdown request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Partners.GetAvailableFunding(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Partners.CreatePartnerData(request) -> *v2.PartnerDataResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Upload supplementary partner information, such as bank account balance, accounting figures, or other relevant details.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.PartnerDataCreatePayload{
        Data: map[string]any{
            "key1": "value1",
            "key2": "value2",
        },
    }
client.Partners.CreatePartnerData(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**data:** `map[string]any` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Partners.ListPartnerWaterfalls() -> *v2.PaginatedResponseWaterfallResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Use this endpoint to get the list of waterfalls for your dedicated lending account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListPartnerWaterfallsRequest{}
client.Partners.ListPartnerWaterfalls(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` — Field to order the results by, e.g., 'created_at:desc,updated_at:asc'
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: id, name, date, status, created_at, updated_at. Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Webhooks
<details><summary><code>client.Webhooks.ListWebhookSubscriptions() -> *v2.PaginatedResponseWebhookSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List all webhook subscriptions for your partner account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListWebhookSubscriptionsRequest{}
client.Webhooks.ListWebhookSubscriptions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**eventType:** `*v2.WebhookEventTypeEnum` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.CreateWebhookSubscription(request) -> *v2.WebhookSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new webhook subscription for a specific event type.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.WebhookCreatePayload{
        URL: "https://example.com/webhooks",
        Description: v2.String(
            "Loan update event",
        ),
        EventType: v2.WebhookEventTypeEnumLoanUpdated,
        Secret: "whsec_f3o9p8h7g6j5k4l3m2n1o0p9i8u7y6t5",
        Retry: v2.Bool(
            false,
        ),
        Status: v2.WebhookStatusEnumActive.Ptr(),
    }
client.Webhooks.CreateWebhookSubscription(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**url:** `string` — The URL to send webhooks to
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — Optional description of this webhook endpoint
    
</dd>
</dl>

<dl>
<dd>

**eventType:** `*v2.WebhookEventTypeEnum` — Event type to subscribe toPossible values: loan_updated, installment_updated, client_updated, client_limit_updated, partner_limit_updated
    
</dd>
</dl>

<dl>
<dd>

**secret:** `string` — Secret for signing webhook payloads
    
</dd>
</dl>

<dl>
<dd>

**retry:** `*bool` — Whether to retry failed webhooks
    
</dd>
</dl>

<dl>
<dd>

**status:** `*v2.WebhookStatusEnum` — Status of the webhook subscription. Defaults to 'active'.Possible values: active, paused, disabled
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.GetWebhookSubscription(WebhookID) -> *v2.WebhookSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details for a specific webhook subscription with its webhook ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.GetWebhookSubscriptionRequest{
        WebhookID: "webhook_id",
    }
client.Webhooks.GetWebhookSubscription(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.UpdateWebhookSubscription(WebhookID, request) -> *v2.WebhookSubscriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a webhook subscription with its specific webhook ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.WebhookUpdatePayload{
        WebhookID: "webhook_id",
        URL: v2.String(
            "https://example.com/webhooks/v2",
        ),
        Description: v2.String(
            "Updated webhook endpoint",
        ),
        EventType: v2.WebhookEventTypeEnumInstallmentUpdated.Ptr(),
        Status: v2.WebhookStatusEnumPaused.Ptr(),
        Retry: v2.Bool(
            true,
        ),
        Secret: v2.String(
            "whsec_updated_secret_here",
        ),
    }
client.Webhooks.UpdateWebhookSubscription(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` — The URL to send webhooks to
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — Description of this webhook endpoint
    
</dd>
</dl>

<dl>
<dd>

**eventType:** `*v2.WebhookEventTypeEnum` — Event type to subscribe toPossible values: loan_updated, installment_updated, client_updated, client_limit_updated, partner_limit_updated
    
</dd>
</dl>

<dl>
<dd>

**status:** `*v2.WebhookStatusEnum` — Status of the webhook subscriptionPossible values: active, paused, disabled
    
</dd>
</dl>

<dl>
<dd>

**retry:** `*bool` — Whether to retry failed webhooks
    
</dd>
</dl>

<dl>
<dd>

**secret:** `*string` — Secret for signing webhook payloads
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.DeleteWebhookSubscription(WebhookID) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific webhook subscription.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.DeleteWebhookSubscriptionRequest{
        WebhookID: "webhook_id",
    }
client.Webhooks.DeleteWebhookSubscription(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.ListWebhookLogs() -> *v2.PaginatedResponseWebhookLogResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all webhook logs linked to your partner account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListWebhookLogsRequest{}
client.Webhooks.ListWebhookLogs(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Repayments
<details><summary><code>client.Repayments.ListRepaymentLogs() -> *v2.PaginatedResponseRepaymentResponseWithClientInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all repayments made under your partner account. Supports filtering by client, loan, or installments.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListRepaymentLogsRequest{}
client.Repayments.ListRepaymentLogs(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**loanID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**installmentID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` — Field to order the results by, e.g., 'created_at:desc,updated_at:asc'
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: id, client_id, loan_id, installment_id, created_at, client.name, client.correlation_id. Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Repayments.CreateRepayment(request) -> *v2.RepaymentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new repayment log for an installment. Requires the installment ID, loan ID or loan correlation ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.RepaymentCreatePayload{
        Amount: &v2.RepaymentCreatePayloadAmount{
            Double: 1.1,
        },
    }
client.Repayments.CreateRepayment(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**installmentID:** `*string` — ID of the installment
    
</dd>
</dl>

<dl>
<dd>

**loanID:** `*string` — ID of the associated Loan
    
</dd>
</dl>

<dl>
<dd>

**correlationID:** `*string` — Correlation ID of associated loan
    
</dd>
</dl>

<dl>
<dd>

**amount:** `*v2.RepaymentCreatePayloadAmount` — The amount of payment made for installment
    
</dd>
</dl>

<dl>
<dd>

**repaymentDate:** `*time.Time` — Please provide the repayment_date if it differs from the time you log this repayment. Otherwise, it will be automatically set.
    
</dd>
</dl>

<dl>
<dd>

**data:** `map[string]any` — Additional metadata related to the repayment
    
</dd>
</dl>

<dl>
<dd>

**isEarlySettlement:** `*bool` — Indicates if this repayment is for early settlement
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Repayments.CreateBulkRepayments(request) -> *v2.BulkRepaymentTaskResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Initiate processing of up to 10000 repayment logs. Returns task_id for tracking progress.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.BulkRepaymentCreatePayload{
        Repayments: []*v2.BulkRepaymentItemPayload{
            &v2.BulkRepaymentItemPayload{
                Amount: &v2.BulkRepaymentItemPayloadAmount{
                    String: "1000.00",
                },
                RepaymentDate: v2.Time(
                    v2.MustParseDateTime(
                        "2023-10-01T12:00:00Z",
                    ),
                ),
                Data: map[string]any{
                    "reference": "TXN-001",
                    "type": "transfer",
                },
                InstallmentID: v2.String(
                    "installment_123",
                ),
            },
            &v2.BulkRepaymentItemPayload{
                Amount: &v2.BulkRepaymentItemPayloadAmount{
                    String: "500.50",
                },
                Data: map[string]any{
                    "notes": "Payment received in office",
                    "type": "cash",
                },
                LoanID: v2.String(
                    "loan_456",
                ),
            },
            &v2.BulkRepaymentItemPayload{
                Amount: &v2.BulkRepaymentItemPayloadAmount{
                    String: "750.00",
                },
                RepaymentDate: v2.Time(
                    v2.MustParseDateTime(
                        "2023-09-30T15:30:00Z",
                    ),
                ),
                CorrelationID: v2.String(
                    "LOAN-789",
                ),
            },
        },
    }
client.Repayments.CreateBulkRepayments(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**repayments:** `[]*v2.BulkRepaymentItemPayload` — List of repayments to create (max 10000)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Repayments.GetBulkRepaymentStatus(TaskID) -> *v2.BulkRepaymentTaskStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Check the progress and results of a bulk repayment processing task.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.GetBulkRepaymentStatusRequest{
        TaskID: "task_id",
    }
client.Repayments.GetBulkRepaymentStatus(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**taskID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Drawdowns
<details><summary><code>client.Drawdowns.ListDrawdowns() -> *v2.PaginatedResponseDrawdownResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of drawdowns.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListDrawdownsRequest{}
client.Drawdowns.ListDrawdowns(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` — Field to order the results by, e.g., 'created_at:desc,updated_at:asc'
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: . Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Drawdowns.CreateDrawdownRequest(request) -> *v2.DrawdownResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new drawdown request.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.DrawdownCreatePayload{
        Amount: &v2.DrawdownCreatePayloadAmount{
            Double: 1.1,
        },
    }
client.Drawdowns.CreateDrawdownRequest(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**amount:** `*v2.DrawdownCreatePayloadAmount` — The amount for the drawdown.
    
</dd>
</dl>

<dl>
<dd>

**drawdownDate:** `*time.Time` — The date for the drawdown. If not provided, defaults to the current date and time.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Drawdowns.ListDrawdownChecklists(DrawdownID) -> *v2.PaginatedResponseDrawdownChecklistResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all checklist items for a specific drawdown
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v2.ListDrawdownChecklistsRequest{
        DrawdownID: "drawdown_id",
    }
client.Drawdowns.ListDrawdownChecklists(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**drawdownID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*string` — Field to order the results by, e.g., 'created_at:desc,updated_at:asc'
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query string for filtering. Format: "field:operator:value;...". Supported fields: . Supported operators: is, in, not_in, contains, not_contains, like, not_like, ilike, not_ilike, gt, gte, lt, lte, starts_with, ends_with, is_null, is_not_null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

