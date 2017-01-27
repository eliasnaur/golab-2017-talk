package org.golang.example.basic;

import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;

import org.golang.example.basic.databinding.ActivityMainBinding;
import android.databinding.DataBindingUtil;

public class MainActivity extends AppCompatActivity {
	@Override public void onCreate(Bundle state) {
		super.onCreate(state);
		ActivityMainBinding binding = DataBindingUtil.setContentView(this, R.layout.activity_main);
	}
}
