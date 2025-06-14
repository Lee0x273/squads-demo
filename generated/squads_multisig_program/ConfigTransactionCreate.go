// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package squads_multisig_program

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Create a new config transaction.
type ConfigTransactionCreate struct {
	Args *ConfigTransactionCreateArgs

	// [0] = [WRITE] multisig
	//
	// [1] = [WRITE] transaction
	//
	// [2] = [SIGNER] creator
	// ··········· The member of the multisig that is creating the transaction.
	//
	// [3] = [WRITE, SIGNER] rentPayer
	// ··········· The payer for the transaction account rent.
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewConfigTransactionCreateInstructionBuilder creates a new `ConfigTransactionCreate` instruction builder.
func NewConfigTransactionCreateInstructionBuilder() *ConfigTransactionCreate {
	nd := &ConfigTransactionCreate{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *ConfigTransactionCreate) SetArgs(args ConfigTransactionCreateArgs) *ConfigTransactionCreate {
	inst.Args = &args
	return inst
}

// SetMultisigAccount sets the "multisig" account.
func (inst *ConfigTransactionCreate) SetMultisigAccount(multisig ag_solanago.PublicKey) *ConfigTransactionCreate {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(multisig).WRITE()
	return inst
}

// GetMultisigAccount gets the "multisig" account.
func (inst *ConfigTransactionCreate) GetMultisigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTransactionAccount sets the "transaction" account.
func (inst *ConfigTransactionCreate) SetTransactionAccount(transaction ag_solanago.PublicKey) *ConfigTransactionCreate {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(transaction).WRITE()
	return inst
}

// GetTransactionAccount gets the "transaction" account.
func (inst *ConfigTransactionCreate) GetTransactionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetCreatorAccount sets the "creator" account.
// The member of the multisig that is creating the transaction.
func (inst *ConfigTransactionCreate) SetCreatorAccount(creator ag_solanago.PublicKey) *ConfigTransactionCreate {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(creator).SIGNER()
	return inst
}

// GetCreatorAccount gets the "creator" account.
// The member of the multisig that is creating the transaction.
func (inst *ConfigTransactionCreate) GetCreatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetRentPayerAccount sets the "rentPayer" account.
// The payer for the transaction account rent.
func (inst *ConfigTransactionCreate) SetRentPayerAccount(rentPayer ag_solanago.PublicKey) *ConfigTransactionCreate {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(rentPayer).WRITE().SIGNER()
	return inst
}

// GetRentPayerAccount gets the "rentPayer" account.
// The payer for the transaction account rent.
func (inst *ConfigTransactionCreate) GetRentPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *ConfigTransactionCreate) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ConfigTransactionCreate {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *ConfigTransactionCreate) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst ConfigTransactionCreate) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ConfigTransactionCreate,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ConfigTransactionCreate) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ConfigTransactionCreate) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Multisig is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Transaction is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Creator is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.RentPayer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *ConfigTransactionCreate) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ConfigTransactionCreate")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     multisig", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  transaction", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      creator", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    rentPayer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj ConfigTransactionCreate) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ConfigTransactionCreate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewConfigTransactionCreateInstruction declares a new ConfigTransactionCreate instruction with the provided parameters and accounts.
func NewConfigTransactionCreateInstruction(
	// Parameters:
	args ConfigTransactionCreateArgs,
	// Accounts:
	multisig ag_solanago.PublicKey,
	transaction ag_solanago.PublicKey,
	creator ag_solanago.PublicKey,
	rentPayer ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *ConfigTransactionCreate {
	return NewConfigTransactionCreateInstructionBuilder().
		SetArgs(args).
		SetMultisigAccount(multisig).
		SetTransactionAccount(transaction).
		SetCreatorAccount(creator).
		SetRentPayerAccount(rentPayer).
		SetSystemProgramAccount(systemProgram)
}
