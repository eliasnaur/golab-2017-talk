package goaccount

import (
	"Java/android/databinding/DataBindingUtil"
	"Java/android/support/v7/widget/LinearLayoutManager"
	"Java/android/support/v7/widget/RecyclerView"
	"Java/android/view"
	"Java/android/view/LayoutInflater"

	gopkg "Java/goaccount"

	rid "Java/goaccount/R/id"
	rlayout "Java/goaccount/R/layout"

	transholder "Java/goaccount/TransactionHolder"

	"Java/goaccount/databinding"
	"Java/goaccount/databinding/ActivityMainBinding"
	"Java/goaccount/databinding/TransactionBinding"
)

// START SETUP OMIT
import transadapter "Java/goaccount/TransactionsAdapter" // HLadapter

func setupActivity(this gopkg.GoActivity) {
	db := DataBindingUtil.SetContentView(this, rlayout.Activity_main) // HLdatabinding
	binding := ActivityMainBinding.Cast(db)                           // HLdatabinding
	rv := RecyclerView.Cast(this.FindViewById(rid.Transactions))      // HLrecyclerview
	rv.SetLayoutManager(LinearLayoutManager.New(this))                // HLrecyclerview
	acc := &Account{}
	acc.NewTransaction(100.00)
	acc.NewTransaction(2.34)
	adapter := transadapter.New()                                           // HLadapter
	adapter.Unwrap().(*TransactionsAdapter).transactions = acc.transactions // HLadapter
	rv.SetAdapter(adapter)
	binding.SetAmount(acc.FormatBalance())
}

// END SETUP OMIT

// START TRANSADAPTER OMIT
type TransactionsAdapter struct {
	RecyclerView.Adapter
	transactions []*Transaction
}

func (a *TransactionsAdapter) GetItemCount() int32 { return int32(len(a.transactions)) }

func (a *TransactionsAdapter) OnBindViewHolder(holder RecyclerView.ViewHolder, pos int32) {
	th := transholder.Cast(holder).Unwrap().(*TransactionHolder)
	t := a.transactions[len(a.transactions)-1-int(pos)]
	th.binding.SetAmount(t.FormatAmount())
	th.binding.SetTime(t.FormatTime())
	th.binding.ExecutePendingBindings()
}

func (a *TransactionsAdapter) OnCreateViewHolder(parent view.ViewGroup, viewType int32) RecyclerView.ViewHolder {
	inflater := LayoutInflater.From(parent.GetContext())
	binding := TransactionBinding.Inflate(inflater, parent, false)
	th := transholder.New(binding.GetRoot())
	th.Unwrap().(*TransactionHolder).binding = binding
	return th
}

// END TRANSADAPTER OMIT

// START TRANSHOLDER OMIT
type TransactionHolder struct {
	RecyclerView.ViewHolder
	binding databinding.TransactionBinding
}

func NewTransactionHolder(_ view.View) *TransactionHolder {
	return new(TransactionHolder)
}

// END TRANSHOLDER OMIT
