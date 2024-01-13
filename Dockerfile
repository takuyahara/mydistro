FROM nixos/nix AS build

RUN echo 'filter-syscalls = false' >> /etc/nix/nix.conf

RUN mkdir -p /output/bun
RUN nix-env --profile /output/profile0 -i bun --filter-syscalls=false
RUN cp -va $(nix-store -qR /output/profile0) /output/bun

RUN mkdir -p /output/store
RUN nix-env --profile /output/profile -i chromium
RUN cp -va $(nix-store -qR /output/profile) /output/store

RUN mkdir -p /output/fonts
RUN nix-env --profile /output/profile2 -i ipafont ipaexfont
RUN cp -va $(nix-store -qR /output/profile2) /output/fonts

RUN ln -s /nix/store/13km4kagqsy1cklxwnw3qyq1apxk86im-chromium-118.0.5993.117/bin/chromium /usr/bin/chromium
RUN ln -s /bun/ws8q22wm8qm41n7fff0xrkzwbjy9l330-bun-0.6.2/bin/bun /usr/bin/bun

# FROM debian
FROM gcr.io/distroless/base-debian12:debug
COPY --from=build /output/bun /bun
COPY --from=build /output/store /nix/store
COPY --from=build /output/fonts /fonts
# COPY --from=build /usr/bin/chromium /usr/bin/chromium
# COPY --from=build /usr/bin/bun /usr/bin/bun
# COPY main /app/main
COPY package.json /app/package.json
COPY main.ts /app/main.ts
# RUN ln -s /nix/store/13km4kagqsy1cklxwnw3qyq1apxk86im-chromium-118.0.5993.117/bin/chromium /usr/bin/chromium
ENTRYPOINT ["deno" "run" "/app/main.ts"]

# cp -r /fonts/6shkjfyk1vb3542i3jva0slqzlk176f5-user-environment/* /usr
# /nix/store/13km4kagqsy1cklxwnw3qyq1apxk86im-chromium-118.0.5993.117/bin/chromium
        # <dir>/usr/share/fonts</dir>