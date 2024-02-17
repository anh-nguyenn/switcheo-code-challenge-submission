import React, { useEffect, useMemo, useState } from 'react';

// Assuming other imports and setup are done correctly

class Datasource {
  private url: string;

  constructor(url: string) {
    this.url = url;
  }

  async getPrices(): Promise<any> {
    try {
      const response = await fetch(this.url);
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return await response.json();
    } catch (error) {
      console.error('There was a problem with your fetch operation:', error);
    }
  }
}

const WalletPage: React.FC<Props> = ({ children, ...rest }) => {
  const balances = useWalletBalances();
  const [prices, setPrices] = useState({});

  useEffect(() => {
    const datasource = new Datasource("https://interview.switcheo.com/prices.json");
    datasource.getPrices().then((prices) => {
      if (prices) {
        setPrices(prices);
      }
    }).catch((error) => {
      console.error(error);
    });
  }, []);

  const getPriority = useMemo(() => {
    const priorityMap = new Map();
    return (blockchain: string): number => {
      if (!priorityMap.has(blockchain)) {
        let priority;
        switch (blockchain) {
          case 'Osmosis': priority = 100; break;
          case 'Ethereum': priority = 50; break;
          case 'Arbitrum': priority = 30; break;
          case 'Zilliqa': priority = 20; break;
          case 'Neo': priority = 20; break;
          default: priority = -99;
        }
        priorityMap.set(blockchain, priority);
      }
      return priorityMap.get(blockchain);
    };
  }, []);

  const sortedBalances = useMemo(() => {
    return balances.filter(balance => balance.amount > 0)
      .sort((lhs, rhs) => {
        const leftPriority = getPriority(lhs.blockchain);
        const rightPriority = getPriority(rhs.blockchain);
        return rightPriority - leftPriority;
      });
  }, [balances, getPriority]);

  const rows = sortedBalances.map((balance, index) => {
    const usdValue = (prices[balance.currency] || 0) * balance.amount;
    return (
      <WalletRow
        key={index}
        amount={balance.amount}
        usdValue={usdValue}
        formattedAmount={balance.amount.toFixed()}
      />
    );
  });

  return <div {...rest}>{rows}</div>;
};
