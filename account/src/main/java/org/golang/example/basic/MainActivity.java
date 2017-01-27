package org.golang.example.basic;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.ViewGroup;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.RecyclerView;
import android.support.v7.widget.LinearLayoutManager;

import org.golang.example.basic.databinding.ActivityMainBinding;
import org.golang.example.basic.databinding.TransactionBinding;
import android.databinding.DataBindingUtil;

import goaccount.Account;

public class MainActivity extends AppCompatActivity {
	@Override public void onCreate(Bundle state) {
		super.onCreate(state);
// START 1 OMIT
		ActivityMainBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_main);
		goaccount.Account account = new goaccount.Account(); // HL
		account.newTransaction(2.34); // HL
		account.newTransaction(100.00); // HL
		binding.setAccount(account); // HL
// END 1 OMIT
		setupTransactions(binding, account);
	}
// START 2 OMIT
	private void setupTransactions(ActivityMainBinding binding, goaccount.Account account) {
		binding.transactions.setLayoutManager(new LinearLayoutManager(this));
		LayoutInflater inflater = LayoutInflater.from(this);
		binding.transactions.setAdapter(new TransactionsAdapter(inflater, account));
	}
// END 2 OMIT
// START 3 OMIT
	private static class TransactionsAdapter extends RecyclerView.Adapter<TransactionHolder> {
		private final goaccount.Account account;
		private final LayoutInflater inflater;

		TransactionsAdapter(LayoutInflater inflater, goaccount.Account account) {
			this.inflater = inflater;
			this.account = account;
		}

		@Override public int getItemCount() {
			return (int)account.numTransactions();
		}

		@Override public void onBindViewHolder(TransactionHolder holder, int position) {
			holder.binding.setTransaction(account.transaction(position));
			holder.binding.executePendingBindings();
		}

		@Override public TransactionHolder onCreateViewHolder(ViewGroup parent, int viewType) {
			TransactionBinding binding = TransactionBinding.inflate(inflater, parent, false);
			return new TransactionHolder(binding);
		}
	}
// END 3 OMIT
// START 4 OMIT
	private static class TransactionHolder extends RecyclerView.ViewHolder {
		TransactionBinding binding;

		TransactionHolder(TransactionBinding binding) {
			super(binding.getRoot());
			this.binding = binding;
		}
	}
// END 4 OMIT
}
