#include "header.h"

// Function to register a new user
void registerMenu(char a[50], char pass[50])
{
    printf("Bla Bla\n [registerMenu]\n");
    //handle registration
}

// Function to update account information
void updateAccountInfo(struct User u)
{
    printf("Bla Bla\n [updateAccountInfo]\n");
    printf("Enter the account number to update: ");
    int accountNumber;
    scanf("%d", &accountNumber);
    printf("Account number you type is: %d\n", accountNumber);
    //check if the account number exists
    // checkAccountNbr(accountNumber) ? printf("Account number exists\n") : printf("Account number does not exist\n");

    //handle account update
}

// Function to check details of a specific account
void checkAccountDetails(struct User u)
{
    printf("Bla Bla\n [checkAccountDetails]\n");
    //handle account details

}

// Function to handle deposits and withdrawals
void makeTransaction(struct User u)
{
    printf("Bla Bla\n [makeTransaction]\n");
    //handle transactions

}

// Function to remove an existing account
void removeAccount(struct User u)
{
    printf("Bla Bla\n [removeAccount]\n");
    //handle account removal

}

void transferOwnership(struct User u)
{
    printf("Bla Bla\n [transferOwnership]\n");
    //handle account transfer

}

// Function to check all accounts of a user
void checkAllAccounts(struct User u)
{
    printf("Bla Bla\n [checkAllAccounts]\n");
    //handle account checking
}