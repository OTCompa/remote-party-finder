use std::collections::HashMap;
use ffxiv_types::jobs::{ClassJob, Class, Job, NonCombatJob};

lazy_static::lazy_static! {
    pub static ref JOBS: HashMap<u32, ClassJob> = maplit::hashmap! {
        1 => ClassJob::Class(Class::Gladiator),
        2 => ClassJob::Class(Class::Pugilist),
        3 => ClassJob::Class(Class::Marauder),
        4 => ClassJob::Class(Class::Lancer),
        5 => ClassJob::Class(Class::Archer),
        6 => ClassJob::Class(Class::Conjurer),
        7 => ClassJob::Class(Class::Thaumaturge),
        8 => ClassJob::NonCombat(NonCombatJob::Carpenter),
        9 => ClassJob::NonCombat(NonCombatJob::Blacksmith),
        10 => ClassJob::NonCombat(NonCombatJob::Armorer),
        11 => ClassJob::NonCombat(NonCombatJob::Goldsmith),
        12 => ClassJob::NonCombat(NonCombatJob::Leatherworker),
        13 => ClassJob::NonCombat(NonCombatJob::Weaver),
        14 => ClassJob::NonCombat(NonCombatJob::Alchemist),
        15 => ClassJob::NonCombat(NonCombatJob::Culinarian),
        16 => ClassJob::NonCombat(NonCombatJob::Miner),
        17 => ClassJob::NonCombat(NonCombatJob::Botanist),
        18 => ClassJob::NonCombat(NonCombatJob::Fisher),
        19 => ClassJob::Job(Job::Paladin),
        20 => ClassJob::Job(Job::Monk),
        21 => ClassJob::Job(Job::Warrior),
        22 => ClassJob::Job(Job::Dragoon),
        23 => ClassJob::Job(Job::Bard),
        24 => ClassJob::Job(Job::WhiteMage),
        25 => ClassJob::Job(Job::BlackMage),
        26 => ClassJob::Class(Class::Arcanist),
        27 => ClassJob::Job(Job::Summoner),
        28 => ClassJob::Job(Job::Scholar),
        29 => ClassJob::Class(Class::Rogue),
        30 => ClassJob::Job(Job::Ninja),
        31 => ClassJob::Job(Job::Machinist),
        32 => ClassJob::Job(Job::DarkKnight),
        33 => ClassJob::Job(Job::Astrologian),
        34 => ClassJob::Job(Job::Samurai),
        35 => ClassJob::Job(Job::RedMage),
        36 => ClassJob::Job(Job::BlueMage),
        37 => ClassJob::Job(Job::Gunbreaker),
        38 => ClassJob::Job(Job::Dancer),
    };
}