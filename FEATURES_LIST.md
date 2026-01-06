# Features List

## General

* I want to have a link to a feature for each step in financial planning. For example, the insurance link should direct users to a page where they can learn some basics about insurance and run some calculations to estimate how much life insurance they need (I want to provide different options for estimating life insurance needs, e.g. DIME, whatever other options there are), does a PPO plan or an High Deductible Plan + an HSA make more sense in their situation, etc.

## Transactions

* Categorizing transactions: The default category will be "Uncategorized". When users click the dropdown button to categorize a transaction, then a modal will popup that will display all of their budget category groups and categories in the same way that they are listed in their budget.
* Reimbursements & Similar Credits: Users will be able to categorize reimbursements as a "Reimbursement" (which is a credit) and they will also be able to credit the amount to a specific expense category. That way the budget total amounts will balance out _AND_ the individual category amounts will also balance out. If users could not credit the amount to an expense category, then only the total amounts would balance out but it would show that the user had spent more in an individual category than they actually did. If the user does not credit the amount to an expense category, then it will only be credited to the income at the top of the budget.
    * To make this work in the budget, I need to have "Budget", "Debit", "Credit", and "Remaining" columns in the budget. I was thinking that I could hide the "Credit" column by default and only show it for a category if a user categorizes a reimbursement for that category.

## Budget

* Users can add category groups and categories directly on the budget page instead of through a separate popup modal.
* Users can sort their budget by: Standard or Date of Expenses. Users can set a due date for each category, which will indicate when the expense is due. Some expenses will not have a due date, like groceries, and those will appear at the end when sorted by date of expenses.

## Financial Check-ins

* "Get finances done without talking about finances." I really do not like having Financial Check-ins, but they are absolutely necessary in order to be on the same page as your spouse and get financial things done. Is there a way to do Financial Check-ins asynchronously, similar to how remote work gets done? In other words, is there a way to work on to-do items when I have time (not when my spouse has time) and report back when it is done? This would give me a greater sense of independence while still working together toward common financial goals when I have the energy and can set aside time to do so (not when my spouse has energy and can set aside time). For example, in remote work a co-worker might DM you, but you don't have to answer until you have time and are able to address the issue. Can we use a to-do list that allows us to track our conversations, notes, and what needs to be done for each to-do item (similar to Jira)? We could assign a to-do item to one of the spouses. Maybe we could schedule a time to work on that to-do item in a calendar app or to receive a reminder notification to work on that item. See how the Church handles calendar reminders after you schedule a temple appointment.
