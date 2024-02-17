/_
Problem 3: Messy React
Author: Anh Nguyen
Date: 16/02/2024
_/

Computational Inefficiencies and Anti-Patterns

1. Repeatedly Calculating Priority: The function getPriority is called multiple times for the same blockchain during sorting. This is inefficient because it recalculates the priority for each comparison without any need. Caching these values could improve efficiency.

2. Filtering Logic in sortedBalances: The filtering condition seems unnecessarily complex. The if (lhsPriority > -99) condition uses an undefined variable lhsPriority. It seems like it should check balancePriority, and the logic itself seems flawed because it returns true for balances with an amount less than or equal to zero, which contradicts the typical intention to filter out such balances.

3. In the useEffect hook, the empty dependency array ([]) means the code runs only once after the first render. If we want to re-fetch data whenever certain component properties change, we should list those properties in the dependency array. This tells React to run the code again whenever those specific properties change, ensuring our data is updated accordingly.

4. Error Handling: The use of console.err is incorrect; it should be console.error.

5. When we start using the prices state by setting it up as an empty object {}, we don't specify what kind of information it will hold. This means we're not telling our code or anyone else looking at it what should be inside prices. For example, it's not clear if prices should have currency names as keys and their values as numbers, or if it should be structured differently. Not being clear about this can cause mistakes or confusion when we try to use prices later in our code

6. Unnecessary Mapping for formattedBalances: The mapping to create formattedBalances is redundant because it formats balances that are already being formatted individually in the rows mapping. This can be simplified.

7. Inefficient Sorting and Filtering: The sorting and filtering logic in useMemo is executed every time either balances or prices changes. Since prices does not seem to directly affect the sorting/filtering logic, its presence in the dependency array may lead to unnecessary recalculations.

8. The useMemo hook in React needs to list all the data it depends on in its dependency array. This ensures it recalculates only when necessary. In the given scenario, it's suggested that useMemo should mainly depend on balances. If any other data affects the memoized value but isn't listed, like prices, React won't know to update when those change, leading to potential errors or outdated information.

Improvement in refactored code:

- Correcting the error handling method to console.error.
- Removing unnecessary mapping and redundancy.
- Improving efficiency in sorting and filtering.
- Ensuring type safety and clarity in state management and function calls.
- Utilizing useMemo for caching and efficiency improvements.
