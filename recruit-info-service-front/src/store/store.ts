import create from "zustand"

type EditedCompany = {
  id: number
  name: string
  description: string
  openSalary: string
}

type State = {
  editedCompany: EditedCompany
  updateEditedCompany: (payload: EditedCompany) => void
  resetEditedCompany: () => void
}

const useStore = create<State>((set) => ({
  editedCompany: { id: 0, name: '', description: '', openSalary: ''},
  updateEditedCompany: (payload) =>
    set({
      editedCompany: payload,
    }),
  resetEditedCompany: () => set({ editedCompany: { id: 0, name: '', description: '', openSalary: '' } }),
}))

export default useStore