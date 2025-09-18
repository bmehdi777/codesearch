import { Input } from "@/components/ui/input";
import { useEffect, useRef, useState } from "react";

interface FormElements extends HTMLFormControlsCollection {
  searchInput: HTMLInputElement;
}
interface SearchFormElement extends HTMLFormElement {
  readonly elements: FormElements;
}

function Search() {
  const [hasSearched, setHasSearched] = useState<boolean>(false);
  const searchInputRef = useRef<HTMLInputElement | null>(null);

  useEffect(() => {
    const handleShortcuts = (event: KeyboardEvent) => {
      if (event.ctrlKey && event.key === 'k') {
				event.preventDefault();
        setHasSearched(false);
        if (searchInputRef.current) {
          searchInputRef.current.focus();
					searchInputRef.current.value = "";
        }
      }
    };

    window.addEventListener("keydown", handleShortcuts);

    return () => {
      window.removeEventListener("keydown", handleShortcuts);
    };
  });

  const handleSubmit = (e: React.FormEvent<SearchFormElement>) => {
    e.preventDefault();
    const searchValue = e.currentTarget.elements.searchInput.value;
    if (searchValue !== "") {
      setHasSearched(true);
    } else {
      setHasSearched(false);
    }
  };

  return (
    <div
      className={`flex flex-col justify-start w-3xl transform duration-100 ease-in-out ${hasSearched ? "justify-start mt-5 mx-auto" : "mt-64 mx-auto"}`}
    >
      <h1 className="text-4xl text-balance text-center mb-8">Codesearch</h1>

      <form onSubmit={handleSubmit}>
        <Input
          ref={searchInputRef}
          name="searchInput"
          type="text"
          placeholder="Search your code..."
        />
      </form>
    </div>
  );
}

export default Search;
