<div align="center">

# DeepFaceKey

[![Build](https://github.com/D4n13l3k00/DeepFaceKey/actions/workflows/build.yml/badge.svg)](https://github.com/D4n13l3k00/DeepFaceKey/actions/workflows/build.yml)
![GitHub](https://img.shields.io/github/license/D4n13l3k00/DeepFaceKey)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/D4n13l3k00/DeepFaceKey)
</div>

## DeepFaceKey - Утилита для удобной загрузки DeepFaceLab на Playkey

### ❓ Зачем

- В [Playkey](https://playkey.net) появился режим "Виртуальный компьютер", позволяющий ставить любой софт и запускать его.
- Эта программа позволяет быстро скачать DeepFaceLab по BitTorrent протоколу (в браузере загрузка будет медленнее)
- На серверах Playkey используются NVIDIA Tesla T4, что улучшает и ускоряет работу нейросети (Лично мне только она и попадатся в вирт. пк)
  - Количество тензорных ядер TURING: 320
  - Количество ядер CUDA: 2560
  - Объем видеопамяти: 16 ГБ GDDR6
  - Пропускная способность памяти: 320+ Гбит/с или выше
  - Производительность операций с одинарной точностью (FP32): 8,1 Терафлопс
  - Производительность операций со смешанной точностью (ML) (FP16/FP32): 65 Терафлопс
  - Производительность операций INT8: 130 тера-операций в секунду (TOPS)
  - Производительность операций INT4: 260 тера-операций в секунду (TOPS)
- Выгодно купить 1 месяц безлимита и тренировать сколько угодно времени (правилами не запрещено, playkey писали что можно запускать софт по типу Photoshop, Premiere, Blender и т.д.)

### 😢 Минусы

- Playkey кикает через пол-часа за афк, в связи с чем нужно часто что-либо делать для продолжения работы машины

### 📜 TODO

- [X] Скачивание DeepFaceLab по BitTorrent протоколу
- [ ] Упаковка проекта в архив и его распаковка
