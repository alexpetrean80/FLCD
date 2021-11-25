from dataclasses import dataclass

@dataclass
class Config:
    state: str
    symbol: str

@dataclass
class FiniteAutomata:
    init_state: str
    states: dict
    alphabet: dict
    final_states: dict
    transitions: dict
    
    def next_state(self, c: Config) -> str:
        states = self.transitions[c]
        is_state_valid = self.alphabet[c.state]
        is_sym_valid = self.alphabet[c.symbol]
        
        if not is_state_valid or not is_sym_valid:
            raise "invalid config"
        
        if len(states) > 1:
            raise "can't compute next state"

        return states[0]

    def is_accepted(self, seq: list) -> bool:
        conf = Config()
        state = self.init_state

        for symbol in seq:
            conf.state = state
            conf.symbol = symbol
            state = self.next_state(conf)

    def is_dererministic(self) -> bool:
        for states in self.transitions:
            if len(states) > 1: 
                return False
        return True


    def add_state(self, new_state: str): 
        self.states[new_state] = True

    def add_symbol(self, new_sym: str):
        self.alphabet[new_sym] = True
                      
    def add_final_state(self, final_state: str):
        self.states[final_state] = True
        self.final_states[final_state] = True
    
    def add_transition(self, state: str, sym: str, next_state: str):
        conf = Config
        conf.state = state
        conf.symbol = sym

        states = self.transitions[conf]

        if states == []:
            self.transitions[conf] = [next_state]

        for s in states:
            if s == next_state:
                return
        
        self.transitions[conf].append(next_state)
