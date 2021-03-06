package bankserv

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/microtest"
	"testing"
)

func TestService_GetUserBankAccounts(t *testing.T) {
	tt := []struct {
		name         string
		exchange     *microtest.Exchange
		bankAccounts BankAccounts
		e            dutil.Error
	}{
		{
			name: "permission required",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			bankAccounts: BankAccounts{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "user not found",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body:   `{"message":"NotFound: Unable to find resource","data":{},"errors":{"user":["not found"]}}`,
				},
			},
			bankAccounts: BankAccounts{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"user": {"not found"},
				},
			},
		},
		{
			name: "user has no bank accounts",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"user bank accounts found","data":{"bank_accounts":[]},"errors":{}}`,
				},
			},
			bankAccounts: BankAccounts{},
			e:            nil,
		},
		{
			name: "user has bank accounts",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"user bank accounts found","data":{"bank_accounts":[{"uuid":"318b052a-7911-4e09-a76d-f6e6a18c6fcd","user_uuid":"ef50ad5f-539a-454d-bb49-c2e3123eaba8","organisation_uuid":null,"account_number":"012345678911","active":true,"create_date":"2022-05-17T04:35:23.000Z","update_date":"2022-05-17T04:35:23.000Z"},{"uuid":"d25ac3b1-0a8f-43a3-8da1-d2f22a814a82","user_uuid":"ef50ad5f-539a-454d-bb49-c2e3123eaba8","organisation_uuid":null,"account_number":"012345678912","active":true,"create_date":"2022-05-17T06:53:32.000Z","update_date":"2022-05-17T06:53:32.000Z"}]},"errors":{}}`,
				},
			},
			bankAccounts: BankAccounts{
				BankAccount{
					UUID:             uuid.MustParse("318b052a-7911-4e09-a76d-f6e6a18c6fcd"),
					UserUUID:         uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8"),
					OrganisationUUID: uuid.UUID{},
					AccountNumber:    "012345678911",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-17T04:35:23.000Z"),
					UpdateDate:       timeMustParse("2022-05-17T04:35:23.000Z"),
				},
				BankAccount{
					UUID:             uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
					UserUUID:         uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8"),
					OrganisationUUID: uuid.UUID{},
					AccountNumber:    "012345678912",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-17T06:53:32.000Z"),
					UpdateDate:       timeMustParse("2022-05-17T06:53:32.000Z"),
				},
			},
			e: nil,
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		UUID := uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8")
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)
			xba, e := s.GetUserBankAccounts(UUID)
			// test that the errors are equal
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			// test the bank accounts
			if len(xba) != len(tc.bankAccounts) {
				t.Errorf(
					"expected bank accounts to have length %d got %d",
					len(tc.bankAccounts),
					len(xba),
				)
			}
			// to check that bank account are equal
			if len(xba) > 0 && tc.bankAccounts[0] != xba[0] {
				t.Errorf(
					"expected bank account %v got %v",
					tc.bankAccounts[0],
					xba[0],
				)
			}
		})
	}
}

func TestService_GetOrganisationBankAccounts(t *testing.T) {
	tt := []struct {
		name         string
		exchange     *microtest.Exchange
		bankAccounts BankAccounts
		e            dutil.Error
	}{
		{
			name: "permission required",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			bankAccounts: BankAccounts{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "organisation not found",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body:   `{"message":"NotFound: Unable to find resource","data":{},"errors":{"organisation":["not found"]}}`,
				},
			},
			bankAccounts: BankAccounts{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"organisation": {"not found"},
				},
			},
		},
		{
			name: "organisation has no bank accounts",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"organisation bank accounts found","data":{"bank_accounts":[]},"errors":{}}`,
				},
			},
			bankAccounts: BankAccounts{},
			e:            nil,
		},
		{
			name: "organisation has bank accounts",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"organisation bank accounts found","data":{"bank_accounts":[{"uuid":"318b052a-7911-4e09-a76d-f6e6a18c6fcd","organisation_uuid":"ef50ad5f-539a-454d-bb49-c2e3123eaba8","user_uuid":null,"account_number":"012345678911","active":true,"create_date":"2022-05-17T04:35:23.000Z","update_date":"2022-05-17T04:35:23.000Z"},{"uuid":"d25ac3b1-0a8f-43a3-8da1-d2f22a814a82","organisation_uuid":"ef50ad5f-539a-454d-bb49-c2e3123eaba8","user_uuid":null,"account_number":"012345678912","active":true,"create_date":"2022-05-17T06:53:32.000Z","update_date":"2022-05-17T06:53:32.000Z"}]},"errors":{}}`,
				},
			},
			bankAccounts: BankAccounts{
				BankAccount{
					UUID:             uuid.MustParse("318b052a-7911-4e09-a76d-f6e6a18c6fcd"),
					UserUUID:         uuid.UUID{},
					OrganisationUUID: uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8"),
					AccountNumber:    "012345678911",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-17T04:35:23.000Z"),
					UpdateDate:       timeMustParse("2022-05-17T04:35:23.000Z"),
				},
				BankAccount{
					UUID:             uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
					UserUUID:         uuid.UUID{},
					OrganisationUUID: uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8"),
					AccountNumber:    "012345678912",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-17T06:53:32.000Z"),
					UpdateDate:       timeMustParse("2022-05-17T06:53:32.000Z"),
				},
			},
			e: nil,
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		UUID := uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8")

		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			xba, e := s.GetOrganisationBankAccounts(UUID)
			// test that the errors are equal
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			// test the bank accounts
			if len(xba) != len(tc.bankAccounts) {
				t.Errorf(
					"expected bank accounts to have length %d got %d",
					len(tc.bankAccounts),
					len(xba),
				)
			}
			// to check that bank account are equal
			if len(xba) > 0 && tc.bankAccounts[0] != xba[0] {
				t.Errorf(
					"expected bank account %v got %v",
					tc.bankAccounts[0],
					xba[0],
				)
			}
		})
	}
}

var userBankAccount = BankAccount{
	UUID:             uuid.MustParse("e6b7f986-307c-4147-a34e-f924790799bb"),
	UserUUID:         uuid.MustParse("e4bd194d-41e7-4f27-a4a8-161685a9b8b8"),
	OrganisationUUID: uuid.UUID{},
	AccountNumber:    "098765432109",
}
var organisationBankAccount = BankAccount{
	UUID:             uuid.MustParse("e6b7f986-307c-4147-a34e-f924790799bb"),
	OrganisationUUID: uuid.MustParse("e4bd194d-41e7-4f27-a4a8-161685a9b8b8"),
	UserUUID:         uuid.UUID{},
	AccountNumber:    "098765432109",
}

func TestService_CreateBankAccount(t *testing.T) {
	tt := []struct {
		name         string
		bankAccount  BankAccount // payload data
		exchange     *microtest.Exchange
		eBankAccount BankAccount // expected bank account
		e            dutil.Error
	}{
		{
			name:        "permission required",
			bankAccount: userBankAccount,
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			eBankAccount: BankAccount{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name:        "user not found",
			bankAccount: userBankAccount,
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body:   `{"message":"NotFound: unable to find resource","data":{},"errors":{"user":["not found"]}}`,
				},
			},
			eBankAccount: BankAccount{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"user": {"not found"},
				},
			},
		},
		{
			name:        "organisation not found",
			bankAccount: organisationBankAccount,
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body:   `{"message":"NotFound: unable to find resource","data":{},"errors":{"organisation":["not found"]}}`,
				},
			},
			eBankAccount: BankAccount{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"organisation": {"not found"},
				},
			},
		},
		{
			name:        "create user bank account",
			bankAccount: userBankAccount,
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 201,
					Body:   `{"message":"bank account create","data":{"bank_account":{"uuid":"e6b7f986-307c-4147-a34e-f924790799bb","user_uuid":"e4bd194d-41e7-4f27-a4a8-161685a9b8b8","organisation_uuid":null,"account_number":"098765432109","active":true,"create_date":"2022-06-17T21:57:12.000Z","update_date":"2022-06-17T21:57:12.000Z"}},"errors":{}}`,
				},
			},
			eBankAccount: BankAccount{
				UUID:             uuid.MustParse("e6b7f986-307c-4147-a34e-f924790799bb"),
				UserUUID:         uuid.MustParse("e4bd194d-41e7-4f27-a4a8-161685a9b8b8"),
				OrganisationUUID: uuid.UUID{},
				AccountNumber:    "098765432109",
				Active:           true,
				CreateDate:       timeMustParse("2022-06-17T21:57:12.000Z"),
				UpdateDate:       timeMustParse("2022-06-17T21:57:12.000Z"),
			},
			e: nil,
		},
		{
			name:        "create organisation bank account",
			bankAccount: organisationBankAccount,
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 201,
					Body:   `{"message":"bank account create","data":{"bank_account":{"uuid":"e6b7f986-307c-4147-a34e-f924790799bb","user_uuid":null,"organisation_uuid":"e4bd194d-41e7-4f27-a4a8-161685a9b8b8","account_number":"098765432109","active":true,"create_date":"2022-06-17T21:57:12.000Z","update_date":"2022-06-17T21:57:12.000Z"}},"errors":{}}`,
				},
			},
			eBankAccount: BankAccount{
				UUID:             uuid.MustParse("e6b7f986-307c-4147-a34e-f924790799bb"),
				UserUUID:         uuid.UUID{},
				OrganisationUUID: uuid.MustParse("e4bd194d-41e7-4f27-a4a8-161685a9b8b8"),
				AccountNumber:    "098765432109",
				Active:           true,
				CreateDate:       timeMustParse("2022-06-17T21:57:12.000Z"),
				UpdateDate:       timeMustParse("2022-06-17T21:57:12.000Z"),
			},
			e: nil,
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			b, e := s.CreateBankAccount(tc.bankAccount)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if b != tc.eBankAccount {
				t.Errorf("expected bank account %v got %v", tc.eBankAccount, b)
			}
		})
	}
}

func TestService_UpdateBankAccount(t *testing.T) {
	tt := []struct {
		name         string
		bankAccount  BankAccount // payload data
		exchange     *microtest.Exchange
		eBankAccount BankAccount // expected bank account
		e            dutil.Error
	}{
		{
			name:        "permission required",
			bankAccount: userBankAccount,
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			eBankAccount: BankAccount{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name:        "update bank account",
			bankAccount: organisationBankAccount,
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"bank account create","data":{"bank_account":{"uuid":"e6b7f986-307c-4147-a34e-f924790799bb","user_uuid":null,"organisation_uuid":"e4bd194d-41e7-4f27-a4a8-161685a9b8b8","account_number":"098765432109","active":true,"create_date":"2022-06-17T22:16:12.000Z","update_date":"2022-06-17T22:16:12.000Z"}},"errors":{}}`,
				},
			},
			eBankAccount: BankAccount{
				UUID:             uuid.MustParse("e6b7f986-307c-4147-a34e-f924790799bb"),
				OrganisationUUID: uuid.MustParse("e4bd194d-41e7-4f27-a4a8-161685a9b8b8"),
				UserUUID:         uuid.UUID{},
				AccountNumber:    "098765432109",
				Active:           true,
				CreateDate:       timeMustParse("2022-06-17T22:16:12.000Z"),
				UpdateDate:       timeMustParse("2022-06-17T22:16:12.000Z"),
			},
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			b, e := s.UpdateBankAccount(tc.bankAccount)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if b != tc.eBankAccount {
				t.Errorf("expected bank account %v got %v", tc.eBankAccount, b)
			}
		})
	}
}

func TestService_DeleteBankAccount(t *testing.T) {
	tt := []struct {
		name     string
		exchange *microtest.Exchange
		UUID     uuid.UUID
		e        dutil.Error
	}{
		{
			name: "permission required",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			e: &dutil.Err{
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "delete bank account",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"bank account deleted","data":{},"errors":{}}`,
				},
			},
			e: nil,
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		UUID := uuid.MustParse("e6b7f986-307c-4147-a34e-f924790799bb")
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			e := s.DeleteBankAccount(UUID)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
		})
	}
}
